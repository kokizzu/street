package domain

import (
	"crypto/rsa"
	"encoding/json"
	"fmt"
	"strings"
	"io/ioutil"
	"math/big"
	"net/http"
	"net/url"
	"time"
	"street/conf"
	"log"
	"encoding/base64"

	"github.com/golang-jwt/jwt/v5"
)

// AppleKey represents a single public key from Apple's JWKS
type AppleKey struct {
	Kty string `json:"kty"`
	Kid string `json:"kid"`
	Use string `json:"use"`
	Alg string `json:"alg"`
	N   string `json:"n"`
	E   string `json:"e"`
}

// AppleKeys represents the response containing multiple public keys from Apple
type AppleKeys struct {
	Keys []AppleKey `json:"keys"`
}

// FetchApplePublicKeys fetches the public keys from Apple's endpoint
func fetchApplePublicKeys() ([]AppleKey, error) {
	resp, err := http.Get("https://appleid.apple.com/auth/keys")
	if err != nil {
		return nil, fmt.Errorf("failed to fetch Apple public keys: %v", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read Apple public keys response: %v", err)
	}

	var keys AppleKeys
	err = json.Unmarshal(body, &keys)
	if err != nil {
		return nil, fmt.Errorf("failed to parse Apple public keys: %v", err)
	}

	return keys.Keys, nil
}

// FindApplePublicKeyByKeyID finds the correct public key based on the key ID (kid)
func findApplePublicKeyByKeyID(keys []AppleKey, kid string) (*AppleKey, error) {
	for _, key := range keys {
		if key.Kid == kid {
			return &key, nil
		}
	}
	return nil, fmt.Errorf("no matching public key found for kid: %s", kid)
}

func decodeSegment(segment string) ([]byte, error) {
	// Base64url decoding (URL safe, no padding)
	segment = strings.TrimRight(segment, "=")
	if l := len(segment) % 4; l > 0 {
		segment += strings.Repeat("=", 4-l)
	}
	return base64.URLEncoding.DecodeString(segment)
}


// ConvertJWKToPublicKey converts an AppleKey (JWK) to an RSA public key
func convertJWKToPublicKey(appleKey *AppleKey) (*rsa.PublicKey, error) {
	// Decode N and E from base64 to big.Int
	nBytes, err := decodeSegment(appleKey.N)
	if err != nil {
		return nil, fmt.Errorf("failed to decode 'n' from JWK: %v", err)
	}
	eBytes, err := decodeSegment(appleKey.E)
	if err != nil {
		return nil, fmt.Errorf("failed to decode 'e' from JWK: %v", err)
	}

	// Convert eBytes to an integer (big-endian)
	e := 0
	for _, b := range eBytes {
		e = e*256 + int(b)
	}

	pubKey := &rsa.PublicKey{
		N: new(big.Int).SetBytes(nBytes),
		E: e,
	}

	return pubKey, nil
}

func exchangeAppleAuthCodeForToken(authCode string, config *conf.AppleOAuthConfig) (accessToken string, idToken string, err error) {
    // Create a JWT for the token request
    token := jwt.NewWithClaims(jwt.SigningMethodES256, jwt.MapClaims{
        "iss": config.TeamID,
        "iat": time.Now().Unix(),
        "exp": time.Now().Add(time.Minute * 5).Unix(),
        "aud": "https://appleid.apple.com",
        "sub": config.ClientID,	  
    })
	token.Header["kid"] = "8KGNT7ZA6F"

	log.Println("authCode => ", authCode)
	log.Println("[exchangeAppleAuthCodeForToken] privateKey => ", config.PrivateKey)
    privateKey, err := jwt.ParseECPrivateKeyFromPEM([]byte(config.PrivateKey))
    if err != nil {
        return "", "", err
    }

	log.Println("PrivateKey => ", privateKey)

    tokenString, err := token.SignedString(privateKey)
    if err != nil {
        return "", "", err
    }

	log.Println("tokenString => ", tokenString)
	log.Println("config.ClientID => ", config.ClientID)

    // Exchange code for access token
    resp, err := http.PostForm("https://appleid.apple.com/auth/token", url.Values{
        "grant_type":    {"authorization_code"},
        "code":          {authCode},
        "client_id":     {config.ClientID},
        "client_secret": {tokenString},
    })

    if err != nil {
		log.Fatalf("HTTP request failed: %v", err)
        return "", "", err
    }
    defer resp.Body.Close()

    var result map[string]interface{}
    if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		log.Fatalf("Failed to decode JSON response: %v", err)
	}
	// Log the entire response
    log.Printf("Response from Apple: %+v", result)

	if errMsg, ok := result["error"].(string); ok {
		log.Fatalf("Apple API returned an error: %s", errMsg)
	}
	
	if errDesc, ok := result["error_description"].(string); ok {
		log.Printf("Error description: %s", errDesc)
	}


    // Safely handle the result map
	accessToken, ok := result["access_token"].(string)
	if !ok || accessToken == "" {
		log.Fatalf("access_token is missing or not a string")
	}
    idToken, ok = result["id_token"].(string)
    if !ok || idToken == "" {
        log.Fatalf("id_token is missing or not a string")
    }

	log.Println("accessToken => ", string(accessToken))
	log.Println("idToken => ", string(idToken))

    return accessToken, idToken, nil
}

func getRSAPublicKey(nStr, eStr string) (*rsa.PublicKey, error) {
	nBytes, err := base64.RawURLEncoding.DecodeString(nStr)
	if err != nil {
		return nil, fmt.Errorf("failed to decode modulus (n): %v", err)
	}
	eBytes, err := base64.RawURLEncoding.DecodeString(eStr)
	if err != nil {
		return nil, fmt.Errorf("failed to decode exponent (e): %v", err)
	}

	n := new(big.Int).SetBytes(nBytes)
	e := int(new(big.Int).SetBytes(eBytes).Uint64())

	return &rsa.PublicKey{
		N: n,
		E: e,
	}, nil
}

func validateAppleIDToken(idToken string) (map[string]interface{}, error) {
	token, err := jwt.Parse(idToken, func(token *jwt.Token) (interface{}, error) {
		// Fetch Apple's public keys and use them to verify the token
		keys, err := fetchApplePublicKeys()
		if err != nil {
			return nil, err
		}

		// Find the key that matches the token's `kid`
		for _, key := range keys {
			if key.Kid == token.Header["kid"] {
				// Convert modulus and exponent to *rsa.PublicKey
				return getRSAPublicKey(key.N, key.E)
			}
		}

		return nil, fmt.Errorf("no matching key found")
	})

	log.Println("[validateAppleIDToken] token:", token)

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}
	log.Println("Invalid token")

	return nil, fmt.Errorf("invalid token")
}


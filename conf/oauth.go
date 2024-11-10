package conf

import (
	"log"
	"os"

	"github.com/kokizzu/gotro/L"
	"github.com/kokizzu/gotro/S"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

// Define the structure for Apple OAuth config
type AppleOAuthConfig struct {
	ClientID    string
	TeamID      string
	KeyID       string
	PrivateKey  string // Private key in PEM format
	RedirectURI string
}

type OauthConf struct {
	Urls        []string
	GoogleScope []string
	Google      map[string]*oauth2.Config
	Apple       map[string]*AppleOAuthConfig // Add Apple OAuth configuration
}

func EnvOauth() (res OauthConf) {
	// Initialize URLs and Google settings
	res.Urls = S.Split(os.Getenv(`OAUTH_URLS`), `,`)
	res.Google = map[string]*oauth2.Config{}
	res.GoogleScope = S.Split(os.Getenv(`OAUTH_GOOGLE_SCOPES`), `,`)
	for _, url := range res.Urls {
		res.Google[url] = &oauth2.Config{
			ClientID:     os.Getenv(`OAUTH_GOOGLE_CLIENT_ID`),
			ClientSecret: os.Getenv(`OAUTH_GOOGLE_CLIENT_SECRET`),
			RedirectURL:  url + `/guest/oauthCallback`,
			Scopes:       res.GoogleScope,
			Endpoint:     google.Endpoint,
		}
	}

	_, err := os.ReadFile(os.Getenv(`OAUTH_APPLE_PRIVATE_KEY_PATH`))
	if err == nil {
		// Initialize Apple settings
		res.Apple = map[string]*AppleOAuthConfig{}
		for _, url := range res.Urls {
			// Load private key from file
			privateKeyPath := os.Getenv(`OAUTH_APPLE_PRIVATE_KEY_PATH`)
			log.Println("[Apple] privateKeyPath => ", privateKeyPath)
			privateKey, err := os.ReadFile(privateKeyPath)
			log.Println("[Apple] privateKeyPath => ", privateKey)
			if err != nil {
				log.Fatalf("Failed to read Apple private key from %s: %v", privateKeyPath, err)
			}

			log.Println("OAUTH_APPLE_CLIENT_ID => ", os.Getenv(`OAUTH_APPLE_CLIENT_ID`))
			log.Println("OAUTH_APPLE_TEAM_ID => ", os.Getenv(`OAUTH_APPLE_TEAM_ID`))
			log.Println("OAUTH_APPLE_KEY_ID => ", os.Getenv(`OAUTH_APPLE_KEY_ID`))

			// Initialize Apple OAuth config
			res.Apple[url] = &AppleOAuthConfig{
				ClientID:    os.Getenv(`OAUTH_APPLE_CLIENT_ID`),       
				TeamID:      os.Getenv(`OAUTH_APPLE_TEAM_ID`),         
				KeyID:       os.Getenv(`OAUTH_APPLE_KEY_ID`),          
				PrivateKey:  string(privateKey),             
				RedirectURI: url + `/guest/oauthCallback`,
			}
		}
	} else {
		L.Print(`OAUTH_APPLE_PRIVATE_KEY_PATH not set, skipping AppleOauth initialization`)
	}

	return
}
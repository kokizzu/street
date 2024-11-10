package domain

import (
    "encoding/json"
	"fmt"
	"log"

	"github.com/kokizzu/gotro/L"
	"github.com/kokizzu/gotro/M"
	"github.com/kokizzu/gotro/S"

	"street/model/mAuth/rqAuth"
	"street/model/mAuth/wcAuth"
)

//go:generate gomodifytags -all -add-tags json,form,query,long,msg -transform camelcase --skip-unexported -w -file GuestOauthCallback.go
//go:generate replacer -afterprefix "Id\" form" "Id,string\" form" type GuestOauthCallback.go
//go:generate replacer -afterprefix "json:\"id\"" "json:\"id,string\"" type GuestOauthCallback.go
//go:generate replacer -afterprefix "By\" form" "By,string\" form" type GuestOauthCallback.go
//go:generate farify doublequote --file GuestOauthCallback.go

type (
	GuestOauthCallbackRedirectIn struct {
		RequestCommon
		State       string `json:"state" form:"state" query:"state" long:"state" msg:"state"`
		Code        string `json:"code" form:"code" query:"code" long:"code" msg:"code"`
		AccessToken string `json:"accessToken" form:"accessToken" query:"accessToken" long:"accessToken" msg:"accessToken"`
	}

	GuestOauthCallbackRedirectOut struct {
		ResponseCommon
		OauthUser   M.SX         `json:"oauthUser" form:"oauthUser" query:"oauthUser" long:"oauthUser" msg:"oauthUser"`
		Email       string       `json:"email" form:"email" query:"email" long:"email" msg:"email"`
		CurrentUser rqAuth.Users `json:"currentUser" form:"currentUser" query:"currentUser" long:"currentUser" msg:"currentUser"`
		Provider    string       `json:"provider" form:"provider" query:"provider" long:"provider" msg:"provider"`

		Segments M.SB `json:"segments" form:"segments" query:"segments" long:"segments" msg:"segments"`
	}
)

const (
	GuestOauthCallbackRedirectAction = `guest/oauthCallbackRedirect`

	ErrGuestOauthCallbackRedirectInvalidState           = `invalid csrf state`
	ErrGuestOauthCallbackRedirectInvalidCsrf            = `invalid csrf token`
	ErrGuestOauthCallbackRedirectInvalidUrl             = `invalid url`
	ErrGuestOauthCallbackRedirectFailedExchange         = `failed exchange oauth token`
	ErrGuestOauthCallbackRedirectFailedUserCreation     = `failed user creation from oauth`
	ErrGuestOauthCallbackRedirectFailedUserModification = `failed user modification from oauth`
	ErrGuestOauthCallbackRedirectFailedStoringSession   = `failed storing session from oauth`
)

func (d *Domain) GuestOauthCallbackRedirect(in *GuestOauthCallbackRedirectIn) (out GuestOauthCallbackRedirectOut) {
	defer d.InsertActionLog(&in.RequestCommon, &out.ResponseCommon)
	// csrf := S.RightOf(in.State, `|`)
	// if csrf == `` {
	// 	out.SetError(400, ErrGuestOauthCallbackInvalidState)
	// 	return
	// }

	// L.Print(in.SessionToken)
	// L.Print(csrf)
	// if !S.StartsWith(in.SessionToken, csrf) {
	// 	out.SetError(400, ErrGuestOauthCallbackInvalidCsrf)
	// 	return
	// }

	providerAccessToken := ""
	fullName := ""

	out.Provider = S.LeftOf(in.State, `|`)

	switch out.Provider {
	case OauthGoogle:
		provider := d.Oauth.Google[in.Host]
		if provider == nil {
			L.Print(in.Host)
			out.SetError(400, ErrGuestOauthCallbackRedirectInvalidUrl)
			return
		}

		token, err := provider.Exchange(in.TracerContext, in.Code)
		if L.IsError(err, `google.provider.Exchange`) {
			out.SetError(400, ErrGuestOauthCallbackRedirectFailedExchange)
			return
		}
		providerAccessToken = string(in.Code)

		client := provider.Client(in.TracerContext, token)
		if d.googleUserInfoEndpointCache == `` {
			json := fetchJsonMap(client, `https://accounts.google.com/.well-known/openid-configuration`, &out.ResponseCommon)
			d.googleUserInfoEndpointCache = json.GetStr(`userinfo_endpoint`)
			if out.HasError() {
				return
			}
		}
		out.OauthUser = fetchJsonMap(client, d.googleUserInfoEndpointCache, &out.ResponseCommon)
		/* from google:
		{
			"email":			"",
			"email_verified":	true,
			"family_name":		"",
			"gender":			"",
			"given_name":		"",
			"locale":			"en-GB",
			"name":				"",
			"picture":			"http://",
			"profile":			"http://",
			"sub":				"number"
		} */
		if out.HasError() {
			return
		}
		out.Email = out.OauthUser.GetStr(`email`)
	case OauthApple:
		provider := d.Oauth.Apple[in.Host]
		if provider == nil {
			out.SetError(400, ErrGuestOauthCallbackRedirectInvalidUrl)
			return
		}

		log.Println("In code => ", in.Code)
	
		// Use apple for exchange apple auth code for tokens
		accessToken, idToken, err := exchangeAppleAuthCodeForToken(in.Code, provider)
		log.Println("Access token => ", string(accessToken))
		if err != nil {
			out.SetError(400, ErrGuestOauthCallbackRedirectFailedExchange)
			return
		}
		providerAccessToken = string(idToken)
		log.Println("providerAccessToken => ", string(providerAccessToken))
		log.Println("idToken => ", string(idToken))
	
		claims, err := validateAppleIDToken(idToken)
		if err != nil {
			out.SetError(400, ErrGuestOauthCallbackRedirectFailedExchange)
			return
		}
	
		out.Email = claims["email"].(string)

		firstNameData := ""
		lastNameData := ""
		if name, ok := claims["name"].(map[string]interface{}); ok {
			if firstName, exists := name["firstName"].(string); exists {
				firstNameData = firstName
			}
		
			if lastName, exists := name["lastName"].(string); exists {
				lastNameData = lastName
			}
		}
		fullName = firstNameData + " " + lastNameData;
	}

	user := wcAuth.NewUsersMutator(d.AuthOltp)
	user.Email = S.ValidateEmail(out.Email)
	user.FullName = fullName

	if !user.FindByEmail() {
		// create user if not exists
		user.VerifiedAt = in.UnixNow()
		user.SetGenUniqueUsernameByEmail(user.Email, in.UnixNow())
		user.SetLastLoginAt(in.UnixNow())

		user.SetUpdatedAt(in.UnixNow())
		user.SetCreatedAt(in.UnixNow())
		if !user.DoInsert() {
			out.SetError(500, ErrGuestOauthCallbackRedirectFailedUserCreation)
			return
		}
		out.actor = user.Id
		out.refId = user.Id
	} else {
		out.actor = user.Id
		out.refId = user.Id

		// update verifiedAt if not verified
		if user.VerifiedAt == 0 {
			user.SetVerifiedAt(in.UnixNow())
		}
		user.SetLastLoginAt(in.UnixNow())

		user.SetUpdatedAt(in.UnixNow())
		if !user.DoUpdateById() {
			out.SetError(500, ErrGuestOauthCallbackRedirectFailedUserModification)
			return
		}
	}

	d.ExpireSession(in.SessionToken, &out.ResponseCommon)

	// create new session
	session, sess := d.CreateSession(user.Id, user.Email, in.UserAgent, in.IpAddress)
	if !session.DoInsert() {
		out.SetError(500, ErrGuestOauthCallbackRedirectFailedStoringSession)
		return
	}

    // Construct JSON data with the session token
	dataWithSessionToken := M.SX{
		"sessionToken": providerAccessToken,
	}

	// Convert data to JSON
	jsonSessionToken, err := json.Marshal(dataWithSessionToken)
	if err != nil {
		out.SetError(500, "Failed to create JSON payload")
		return
	}

    // URL encode the JSON payload
	encodedPayload := string(jsonSessionToken)

    // Construct the custom URL scheme
	redirectUrl := fmt.Sprintf("uniwebview://oauth?code=%s", encodedPayload)

	// Redirect to the custom URL scheme
    out.SetRedirect(redirectUrl)


	out.SessionToken = session.SessionToken
	out.Segments = sess.Segments

	return
}

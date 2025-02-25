package auth

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"io"
	"linker/app/services"
	"linker/internal/structs"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
)


var LinkedIn = oauth2.Endpoint{
	AuthURL:  "https://www.linkedin.com/oauth/v2/authorization",
	TokenURL: "https://www.linkedin.com/oauth/v2/accessToken",
}

var linkedInOauthConfig = &oauth2.Config{
	RedirectURL:  "http://localhost:8080/auth/linkedin/callback",
	ClientID:     os.Getenv("LINKEDIN_OAUTH_CLIENT_ID"),
	ClientSecret: os.Getenv("LINKEDIN_OAUTH_CLIENT_SECRET"),
	Scopes:       []string{"openid email profile"},
	Endpoint:     LinkedIn,
}

const oauthLinkedInApiUrl = "https://api.linkedin.com/v2/userinfo"

func OauthLogin(ctx *gin.Context) {
	oauthState := generateStateOauthCookie(ctx)

	authURL := linkedInOauthConfig.AuthCodeURL(oauthState)
    
	ctx.Redirect(http.StatusTemporaryRedirect, authURL)
}

func OauthCallBack(ctx *gin.Context) {
	code := ctx.Query("code")

	oauthState, err := ctx.Cookie("linker_oauthstate")
	if err != nil {
		log.Printf("Error getting cookie: %v \n", err)
	}

	if ctx.Query("state") != oauthState {
        log.Printf("Invalid oauth state: %v \n", oauthState)
		ctx.Redirect(http.StatusPermanentRedirect, "/")
		return
	}

	data, err := getUserData(code)
	if err != nil {
		log.Printf("Error getting user data: %v \n", err)
		ctx.Redirect(http.StatusPermanentRedirect, "/")
		return
	}

    var userData structs.UserDataLinkedIn

	err = json.Unmarshal(data, &userData)
	if err != nil {
		log.Printf("Error unmarshaling user data: %v \n", err)
		return
	}

    user, err := services.LinkedInGetUser(userData)
    if err != nil {
        log.Printf("Error getting user: %v \n", err)
        return 
    }

    tokenString, err := services.GenerateToken(user)
    if err != nil {
        log.Printf("Error generating token: %v \n", err)
    }

    services.SetTokenCookie(ctx, tokenString)
}

func getUserData(code string) ([]byte, error) {
	token, err := linkedInOauthConfig.Exchange(context.Background(), code)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("GET", oauthLinkedInApiUrl, nil)
	if err != nil {
		return nil, err
	}

	request.Header.Set("Authorization", "Bearer "+token.AccessToken)

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != err {
		return nil, err
	}

	return body, nil
}

func generateStateOauthCookie(ctx *gin.Context) string {
	expiration := 60 * 60

	b := make([]byte, 16)
	rand.Read(b)
	state := base64.URLEncoding.EncodeToString(b)
	ctx.SetCookie("linker_oauthstate", state, expiration, "/", "localhost", false, true)

	return state
}


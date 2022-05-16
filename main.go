package main

import (
	"encoding/json"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"

	logger "github.com/rs/zerolog/log"
)

type OrganizationDetails struct {
	Name string `json:"name"`
}

type UserDetails struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type Claims struct {
	AppID               string              `json:"app_id"`
	UserID              string              `json:"user_id"`
	OrganizationID      string              `json:"organization_id"`
	Userdetails         UserDetails         `json:"user_details"`
	Organizationdetails OrganizationDetails `json:"organization_details"`
	jwt.StandardClaims
}

type CordConfig struct {
	AppID  string `json:"app_id"`
	Secret string `json:"secret"`
}

var Config CordConfig

func initCord() {
	data := CordConfig{}
	s := os.Getenv("CORD")
	json.Unmarshal([]byte(s), &data)
	Config = data

	if Config.AppID == "" {
		logger.Fatal().Msg("CORD_APP_ID env is unset")
	}
	if Config.Secret == "" {
		logger.Fatal().Msg("CORD_SECRET env is unset")
	}

}

func main() {
	initCord()
	jwtKey := []byte(Config.Secret)

	// Declare the expiration time of the token
	// here, we have kept it as 1 minute
	expirationTime := time.Now().Add(1 * time.Minute)
	// Create the JWT claims, which includes the username and expiry time
	claims := &Claims{
		AppID:          Config.AppID,
		UserID:         "USER_ID",
		OrganizationID: "ORG_ID",
		Organizationdetails: OrganizationDetails{
			Name: "ORG_NAME",
		},
		Userdetails: UserDetails{
			Name:  "USER_NAME",
			Email: "USER_EMAIL",
		},
		StandardClaims: jwt.StandardClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: expirationTime.Unix(),
		},
	}

	// Declare the token with the algorithm used for signing, and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	// Create the JWT string
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		// If there is an error in creating the JWT return an internal server error
		logger.Fatal().Msg("Error signing the jwtKey")
		return
	}

	logger.Info().Msg(tokenString)
}

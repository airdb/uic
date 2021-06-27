package service

import (
	"fmt"

	"github.com/airdb/uic/internal/app"
	"github.com/airdb/uic/internal/app/domain/valueobject"
)

func GetOauthConfig() valueobject.OauthConfig {
	a := valueobject.OauthConfig{}

	c := app.InitInjection()
	fmt.Println(c)

	a.ID = c.ID
	a.ClientID = c.ClientID
	a.RedirectURL = c.RedirectURL
	a.State = ""

	return a
}

func GetOauthRedirectURL() string {
	config := GetOauthConfig()

	loginURL := fmt.Sprintf("%s?client_id=%s&redirect_uri=%s&response_type=code&scope=%s&state=%s",
		"https://github.com/login/oauth/authorize",
		config.ClientID,
		config.RedirectURL,
		"user:email read:org",
		config.State,
	)

	return loginURL
}

// Bitbank is an bitcoin exchange
type Bitbank struct{}

func GetUser() valueobject.User {
	a := valueobject.User{
		// ID:       1,
		// Username: "dean",
	}

	c := app.InitInjectionUser()

	a.ID = c.ID
	a.Username = c.Username
	a.Token = "68b329da9893e34099c7d8ad5cb9c940"

	fmt.Println(c)

	return a
}

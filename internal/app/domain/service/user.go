package service

import (
	"fmt"

	"github.com/airdb/uic/internal/app"
	"github.com/airdb/uic/internal/app/adapter/repository"
	"github.com/airdb/uic/internal/app/domain/valueobject"
)

// IExchange is interface of bitcoin exchange
type IExchange interface {
	GetUser() valueobject.User
	// Ticker(p valueobject.Pair) valueobject.Ticker
}

type IUser interface {
	Get() repository.User
	Update(repository.User)

	GetOauthConfig() repository.OauthConfig
}

func (b Bitbank) GetUser() valueobject.User {
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

/*
func convertToCandlestick(res bitbankohlcresponse) []valueobject.CandleStick {
	ohlcs := res.Data.Candlestick[0].Ohlcv
	cs := make([]valueobject.CandleStick, 0)
	for _, v := range ohlcs {
		var base1000 float64 = 1000
		base10 := 10
		timestamp := strconv.FormatInt(int64(v[5].(float64)/base1000), base10)
		cs = append(cs, valueobject.CandleStick{
			Open:      v[0].(string),
			High:      v[1].(string),
			Low:       v[2].(string),
			Close:     v[3].(string),
			Volume:    v[4].(string),
			Timestamp: timestamp,
		})
	}
	return cs
}
*/

func GetUser(e IExchange) valueobject.User {
	return e.GetUser()
}

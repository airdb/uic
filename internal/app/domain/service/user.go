package service

import (
	"fmt"

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
}

func (b Bitbank) GetUser() valueobject.User {
	a := valueobject.User{
		// ID:       1,
		// Username: "dean",
	}

	u := repository.User{}
	c := u.Get()

	a.ID = c.ID
	a.Username = c.Username

	fmt.Println(c)

	return a
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

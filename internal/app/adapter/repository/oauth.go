package repository

import (
	"fmt"
)

func GetOauthConfig() OauthConfig {
	fmt.Println("hello hi")

	db := Connection()
	var config OauthConfig

	result := db.Find(&config)
	if result.Error != nil {
		panic(result.Error)
	}

	return config
}

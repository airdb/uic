package repository

import (
	"fmt"

	"github.com/airdb/uic/internal/app/domain"
	"github.com/airdb/uic/internal/app/domain/factory"
	"gorm.io/gorm"
)

// User is the repository of domain.User
type OauthConfig struct {
	gorm.Model
	Provider    string `gorm:"column:provider"`
	ClientID    string `gorm:"column:client_id"`
	CientSecret string `gorm:"column:cient_secret"`
	RedirectURL string `gorm:"column:redirect_url"`
	State       string `gorm:"column:state"`
}

// Get gets order
func (o OauthConfig) GetOauthConfig() domain.OauthConfig {
	db := Connection()
	var config OauthConfig

	result := db.Find(&config)
	if result.Error != nil {
		panic(result.Error)
	}

	orderFactory := factory.OauthConfig{}
	fmt.Println("xxx", config)
	return orderFactory.GenerateOauthConfig(
		config.ID,
		config.ClientID,
		config.RedirectURL,
		config.State,
	)
}

func Hello() OauthConfig {
	fmt.Println("hello hi")
	return OauthConfig{ClientID: "11"}
}

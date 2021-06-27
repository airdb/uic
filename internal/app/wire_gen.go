// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package app

import (
	"github.com/airdb/uic/internal/app/adapter/repository"
)

// Injectors from wire.go:
func InitInjection() repository.OauthConfig {
	oauthConfig := repository.Hello()
	return oauthConfig
}
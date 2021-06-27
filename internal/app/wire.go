//+build wireinject

// The build tag makes sure the stub is not built in the final build.
package app

import (
	"github.com/airdb/uic/internal/app/adapter/repository"
	"github.com/google/wire"
)

// InitializeEvent creates an Event. It will error if the Event is staffed with
// a grumpy greeter.
func InitInjection() repository.OauthConfig {
	wire.Build(repository.Hello)

	return repository.OauthConfig{}
}

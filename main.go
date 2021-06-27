package main

import (
	"github.com/airdb/sailor/version"
	"github.com/airdb/uic/internal/app/adapter"
)

func main() {
	version.InitUptime()

	adapter.NewRouter()
}

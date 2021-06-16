package main

import (
	"time"

	"github.com/airdb/sailor/version"
	"github.com/airdb/uic/internal/app/adapter"
)

func main() {
	// Init the loc.
	loc, _ := time.LoadLocation("Asia/Shanghai")

	// Set timezone.
	version.CreatedAt = time.Now().In(loc)

	adapter.NewRouter()
}

package main

import (
	"fmt"
	"time"

	"github.com/airdb/sailor/version"
	"github.com/airdb/uic/internal/app"
	"github.com/airdb/uic/internal/app/adapter"
)

func main() {
	a := app.InitInjection()
	fmt.Println(a)

	// Init the loc.
	loc, _ := time.LoadLocation("Asia/Shanghai")

	// Set timezone.
	version.CreatedAt = time.Now().In(loc)

	adapter.NewRouter()
}

//go:generate goagen bootstrap -d crypto\design

package main

import (
	"crypto/app"

	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
)

func main() {
	// Create service
	service := goa.New("crypto")

	// Mount middleware
	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest(true))
	service.Use(middleware.ErrorHandler(service, true))
	service.Use(middleware.Recover())

	// Mount "currency" controller
	c := NewCurrencyController(service)
	app.MountCurrencyController(service, c)

	// Start service
	if err := service.ListenAndServe(":80"); err != nil {
		service.LogError("startup", "err", err)
	}

}

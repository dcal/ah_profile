package main

import (
	"github.com/alliancehealth/ah_profile/app"
	ct "github.com/alliancehealth/ah_profile/controller"
	"github.com/alliancehealth/ah_profile/db"
	"github.com/alliancehealth/ah_profile/swagger"
	"github.com/goadesign/goa"
	"github.com/goadesign/middleware"
)

func main() {
	defer db.DB.Close()

	// Create service
	service := goa.New("API")

	// Setup middleware
	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest())
	service.Use(middleware.LogResponse())
	service.Use(middleware.Recover())

	// Mount "users" controller
	c := ct.NewUsersController(service)
	app.MountUsersController(service, c)
	// Mount Swagger spec provider controller
	swagger.MountController(service)

	// Start service, listen on port 8080
	service.ListenAndServe(":8080")
}

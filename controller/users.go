package controller

import (
	"database/sql"
	"log"

	"github.com/alliancehealth/ah_profile/app"
	"github.com/alliancehealth/ah_profile/db"
	"github.com/goadesign/goa"
)

// UsersController implements theusers resource.
type UsersController struct {
	db *sql.DB
	*goa.Controller
}

// NewUsersController creates a users controller.
func NewUsersController(service *goa.Service) app.UsersController {
	data, err := db.Connection()
	if err != nil {
		log.Println(err.Error())
	}
	return &UsersController{Controller: service.NewController("users"), db: data}
}

// Create runs the create action.
func (c *UsersController) Create(ctx *app.CreateUsersContext) error {
	p := ctx.Payload
	m := db.UserModel{
		Active:       p.Active,
		Email:        *p.Email,
		FirstName:    p.FirstName,
		LastName:     p.LastName,
		Phone:        p.Phone,
		Prescription: *p.Prescription,
	}
	m, err := db.CreateUser(m)
	if err != nil {
		return err
	}
	usr := c.toUserMedia(m)
	return ctx.Created(usr)
}

func (c *UsersController) Show(ctx *app.ShowUsersContext) error {
	m := db.UserModel{ID: ctx.UserID}
	m, err := db.ShowUser(m)
	if err != nil {
		return err
	}
	usr := c.toUserMedia(m)
	return ctx.OK(usr)
}

func (c *UsersController) List(ctx *app.ListUsersContext) error {
	uc := []*app.User{}
	usrs, err := db.ListUsers()
	if err != nil {
		return err
	}
	for _, usr := range usrs {
		um := c.toUserMedia(usr)
		uc = append(uc, um)
	}
	return ctx.OK(uc)
}

func (c *UsersController) Delete(ctx *app.DeleteUsersContext) error {
	m := db.UserModel{ID: ctx.UserID}
	m, err := db.DeleteUser(m)
	if err != nil {
		return err
	}
	um := c.toUserMedia(m)
	return ctx.OK(um)

}

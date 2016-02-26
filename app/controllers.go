//************************************************************************//
// API "ah_profile": Application Controllers
//
// Generated with goagen v0.0.1, command line:
// $ goagen
// --out=github.com/alliancehealth/ah_profile
// --design=github.com/alliancehealth/ah_profile/design
// --pkg=app
//
// The content of this file is auto-generated, DO NOT MODIFY
//************************************************************************//

package app

import "github.com/goadesign/goa"

// UsersController is the controller interface for the Users actions.
type UsersController interface {
	goa.Controller
	Create(*CreateUsersContext) error
	Delete(*DeleteUsersContext) error
	List(*ListUsersContext) error
	Show(*ShowUsersContext) error
}

// MountUsersController "mounts" a Users resource controller on the given service.
func MountUsersController(service goa.Service, ctrl UsersController) {
	// Setup encoders and decoders. This is idempotent and is done by each MountXXX function.
	service.SetEncoder(goa.JSONEncoderFactory(), true, "application/json")
	service.SetEncoder(goa.XMLEncoderFactory(), false, "application/xml")
	service.SetDecoder(goa.JSONDecoderFactory(), true, "application/json")
	service.SetDecoder(goa.XMLDecoderFactory(), false, "application/xml")

	// Setup endpoint handler
	var h goa.Handler
	mux := service.ServeMux()
	h = func(c *goa.Context) error {
		ctx, err := NewCreateUsersContext(c)
		if err != nil {
			return goa.NewBadRequestError(err)
		}
		if rawPayload := ctx.RawPayload(); rawPayload != nil {
			ctx.Payload = rawPayload.(*CreateUsersPayload)
		}
		return ctrl.Create(ctx)
	}
	mux.Handle("POST", "/users", ctrl.HandleFunc("Create", h, unmarshalCreateUsersPayload))
	service.Info("mount", "ctrl", "Users", "action", "Create", "route", "POST /users")
	h = func(c *goa.Context) error {
		ctx, err := NewDeleteUsersContext(c)
		if err != nil {
			return goa.NewBadRequestError(err)
		}
		return ctrl.Delete(ctx)
	}
	mux.Handle("DELETE", "/users/:user_id", ctrl.HandleFunc("Delete", h, nil))
	service.Info("mount", "ctrl", "Users", "action", "Delete", "route", "DELETE /users/:user_id")
	h = func(c *goa.Context) error {
		ctx, err := NewListUsersContext(c)
		if err != nil {
			return goa.NewBadRequestError(err)
		}
		return ctrl.List(ctx)
	}
	mux.Handle("GET", "/users", ctrl.HandleFunc("List", h, nil))
	service.Info("mount", "ctrl", "Users", "action", "List", "route", "GET /users")
	h = func(c *goa.Context) error {
		ctx, err := NewShowUsersContext(c)
		if err != nil {
			return goa.NewBadRequestError(err)
		}
		return ctrl.Show(ctx)
	}
	mux.Handle("GET", "/users/:user_id", ctrl.HandleFunc("Show", h, nil))
	service.Info("mount", "ctrl", "Users", "action", "Show", "route", "GET /users/:user_id")
}

// unmarshalCreateUsersPayload unmarshals the request body.
func unmarshalCreateUsersPayload(ctx *goa.Context) error {
	payload := &CreateUsersPayload{}
	if err := ctx.Service().DecodeRequest(ctx, payload); err != nil {
		return err
	}
	if err := payload.Validate(); err != nil {
		return err
	}
	ctx.SetPayload(payload)
	return nil
}

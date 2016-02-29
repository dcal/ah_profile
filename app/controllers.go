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

import (
	"github.com/goadesign/goa"
	"golang.org/x/net/context"
	"net/http"
)

// inited is true if initService has been called
var inited = false

// initService sets up the service encoders, decoders and mux.
func initService(service *goa.Service) {
	if inited {
		return
	}
	inited = true
	// Setup encoders and decoders
	service.SetEncoder(goa.JSONEncoderFactory(), true, "application/json")
	service.SetEncoder(goa.XMLEncoderFactory(), false, "application/xml")
	service.SetDecoder(goa.JSONDecoderFactory(), true, "application/json")
	service.SetDecoder(goa.XMLDecoderFactory(), false, "application/xml")
}

// UsersController is the controller interface for the Users actions.
type UsersController interface {
	goa.Muxer
	Create(*CreateUsersContext) error
	Delete(*DeleteUsersContext) error
	List(*ListUsersContext) error
	Show(*ShowUsersContext) error
}

// MountUsersController "mounts" a Users resource controller on the given service.
func MountUsersController(service *goa.Service, ctrl UsersController) {
	initService(service)
	var h goa.Handler
	mux := service.Mux
	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		rctx, err := NewCreateUsersContext(ctx)
		if err != nil {
			return goa.NewBadRequestError(err)
		}
		if rawPayload := goa.Request(ctx).Payload; rawPayload != nil {
			rctx.Payload = rawPayload.(*CreateUsersPayload)
		}
		return ctrl.Create(rctx)
	}
	mux.Handle("POST", "/users", ctrl.MuxHandler("Create", "", h, unmarshalCreateUsersPayload))
	goa.Info(goa.RootContext, "mount", goa.KV{"ctrl", "Users"}, goa.KV{"action", "Create"}, goa.KV{"route", "POST /users"})
	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		rctx, err := NewDeleteUsersContext(ctx)
		if err != nil {
			return goa.NewBadRequestError(err)
		}
		return ctrl.Delete(rctx)
	}
	mux.Handle("DELETE", "/users/:user_id", ctrl.MuxHandler("Delete", "", h, nil))
	goa.Info(goa.RootContext, "mount", goa.KV{"ctrl", "Users"}, goa.KV{"action", "Delete"}, goa.KV{"route", "DELETE /users/:user_id"})
	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		rctx, err := NewListUsersContext(ctx)
		if err != nil {
			return goa.NewBadRequestError(err)
		}
		return ctrl.List(rctx)
	}
	mux.Handle("GET", "/users", ctrl.MuxHandler("List", "", h, nil))
	goa.Info(goa.RootContext, "mount", goa.KV{"ctrl", "Users"}, goa.KV{"action", "List"}, goa.KV{"route", "GET /users"})
	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		rctx, err := NewShowUsersContext(ctx)
		if err != nil {
			return goa.NewBadRequestError(err)
		}
		return ctrl.Show(rctx)
	}
	mux.Handle("GET", "/users/:user_id", ctrl.MuxHandler("Show", "", h, nil))
	goa.Info(goa.RootContext, "mount", goa.KV{"ctrl", "Users"}, goa.KV{"action", "Show"}, goa.KV{"route", "GET /users/:user_id"})
}

// unmarshalCreateUsersPayload unmarshals the request body into the context request data Payload field.
func unmarshalCreateUsersPayload(ctx context.Context, req *http.Request) error {
	var payload CreateUsersPayload
	if err := goa.RequestService(ctx).DecodeRequest(req, &payload); err != nil {
		return err
	}
	if err := payload.Validate(); err != nil {
		return err
	}
	goa.Request(ctx).Payload = &payload
	return nil
}

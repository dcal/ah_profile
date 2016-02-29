//************************************************************************//
// API "ah_profile": Application Contexts
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
	"strconv"
)

// CreateUsersContext provides the users create action context.
type CreateUsersContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	Payload *CreateUsersPayload
}

// NewCreateUsersContext parses the incoming request URL and body, performs validations and creates the
// context used by the users controller create action.
func NewCreateUsersContext(ctx context.Context) (*CreateUsersContext, error) {
	var err error
	req := goa.Request(ctx)
	rctx := CreateUsersContext{Context: ctx, ResponseData: goa.Response(ctx), RequestData: req}
	return &rctx, err
}

// CreateUsersPayload is the users create action payload.
type CreateUsersPayload struct {
	// 1 = active, 0 = inactive
	Active    string  `json:"active" xml:"active"`
	CreatedAt *string `json:"created_at,omitempty" xml:"created_at,omitempty"`
	// The user's email
	Email *string `json:"email,omitempty" xml:"email,omitempty"`
	// The user's first name
	FirstName string `json:"firstName" xml:"firstName"`
	// The user's last name
	LastName string `json:"lastName" xml:"lastName"`
	// The user's phone number
	Phone string `json:"phone" xml:"phone"`
	// The user's prescription
	Prescription *string `json:"prescription,omitempty" xml:"prescription,omitempty"`
}

// Validate runs the validation rules defined in the design.
func (payload *CreateUsersPayload) Validate() (err error) {
	if payload.FirstName == "" {
		err = goa.MissingAttributeError(`raw`, "firstName", err)
	}
	if payload.LastName == "" {
		err = goa.MissingAttributeError(`raw`, "lastName", err)
	}
	if payload.Phone == "" {
		err = goa.MissingAttributeError(`raw`, "phone", err)
	}
	if payload.Active == "" {
		err = goa.MissingAttributeError(`raw`, "active", err)
	}

	if ok := goa.ValidatePattern(`^(0|1)$`, payload.Active); !ok {
		err = goa.InvalidPatternError(`raw.active`, payload.Active, `^(0|1)$`, err)
	}
	if payload.Email != nil {
		if err2 := goa.ValidateFormat(goa.FormatEmail, *payload.Email); err2 != nil {
			err = goa.InvalidFormatError(`raw.email`, *payload.Email, goa.FormatEmail, err2, err)
		}
	}
	if ok := goa.ValidatePattern(`^\d{10,}$`, payload.Phone); !ok {
		err = goa.InvalidPatternError(`raw.phone`, payload.Phone, `^\d{10,}$`, err)
	}
	return
}

// Created sends a HTTP response with status code 201.
func (ctx *CreateUsersContext) Created(r *User) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.user+json")
	return ctx.ResponseData.Send(ctx.Context, 201, r)
}

// BadRequest sends a HTTP response with status code 400.
func (ctx *CreateUsersContext) BadRequest() error {
	ctx.ResponseData.WriteHeader(400)
	return nil
}

// InternalServerError sends a HTTP response with status code 500.
func (ctx *CreateUsersContext) InternalServerError() error {
	ctx.ResponseData.WriteHeader(500)
	return nil
}

// DeleteUsersContext provides the users delete action context.
type DeleteUsersContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	UserID int
}

// NewDeleteUsersContext parses the incoming request URL and body, performs validations and creates the
// context used by the users controller delete action.
func NewDeleteUsersContext(ctx context.Context) (*DeleteUsersContext, error) {
	var err error
	req := goa.Request(ctx)
	rctx := DeleteUsersContext{Context: ctx, ResponseData: goa.Response(ctx), RequestData: req}
	rawUserID := req.Params.Get("user_id")
	if rawUserID != "" {
		if userID, err2 := strconv.Atoi(rawUserID); err2 == nil {
			rctx.UserID = userID
		} else {
			err = goa.InvalidParamTypeError("user_id", rawUserID, "integer", err)
		}
	}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *DeleteUsersContext) OK(r *User) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.user+json")
	return ctx.ResponseData.Send(ctx.Context, 200, r)
}

// BadRequest sends a HTTP response with status code 400.
func (ctx *DeleteUsersContext) BadRequest() error {
	ctx.ResponseData.WriteHeader(400)
	return nil
}

// InternalServerError sends a HTTP response with status code 500.
func (ctx *DeleteUsersContext) InternalServerError() error {
	ctx.ResponseData.WriteHeader(500)
	return nil
}

// ListUsersContext provides the users list action context.
type ListUsersContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
}

// NewListUsersContext parses the incoming request URL and body, performs validations and creates the
// context used by the users controller list action.
func NewListUsersContext(ctx context.Context) (*ListUsersContext, error) {
	var err error
	req := goa.Request(ctx)
	rctx := ListUsersContext{Context: ctx, ResponseData: goa.Response(ctx), RequestData: req}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *ListUsersContext) OK(r UserCollection) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.user+json; type=collection")
	return ctx.ResponseData.Send(ctx.Context, 200, r)
}

// BadRequest sends a HTTP response with status code 400.
func (ctx *ListUsersContext) BadRequest() error {
	ctx.ResponseData.WriteHeader(400)
	return nil
}

// InternalServerError sends a HTTP response with status code 500.
func (ctx *ListUsersContext) InternalServerError() error {
	ctx.ResponseData.WriteHeader(500)
	return nil
}

// ShowUsersContext provides the users show action context.
type ShowUsersContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	UserID int
}

// NewShowUsersContext parses the incoming request URL and body, performs validations and creates the
// context used by the users controller show action.
func NewShowUsersContext(ctx context.Context) (*ShowUsersContext, error) {
	var err error
	req := goa.Request(ctx)
	rctx := ShowUsersContext{Context: ctx, ResponseData: goa.Response(ctx), RequestData: req}
	rawUserID := req.Params.Get("user_id")
	if rawUserID != "" {
		if userID, err2 := strconv.Atoi(rawUserID); err2 == nil {
			rctx.UserID = userID
		} else {
			err = goa.InvalidParamTypeError("user_id", rawUserID, "integer", err)
		}
	}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *ShowUsersContext) OK(r *User) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.user+json")
	return ctx.ResponseData.Send(ctx.Context, 200, r)
}

// BadRequest sends a HTTP response with status code 400.
func (ctx *ShowUsersContext) BadRequest() error {
	ctx.ResponseData.WriteHeader(400)
	return nil
}

// NotFound sends a HTTP response with status code 404.
func (ctx *ShowUsersContext) NotFound() error {
	ctx.ResponseData.WriteHeader(404)
	return nil
}

// InternalServerError sends a HTTP response with status code 500.
func (ctx *ShowUsersContext) InternalServerError() error {
	ctx.ResponseData.WriteHeader(500)
	return nil
}

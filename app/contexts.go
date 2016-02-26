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

import "github.com/goadesign/goa"

// CreateUsersContext provides the users create action context.
type CreateUsersContext struct {
	*goa.Context
	Payload *CreateUsersPayload
}

// NewCreateUsersContext parses the incoming request URL and body, performs validations and creates the
// context used by the users controller create action.
func NewCreateUsersContext(c *goa.Context) (*CreateUsersContext, error) {
	var err error
	ctx := CreateUsersContext{Context: c}
	return &ctx, err
}

// CreateUsersPayload is the users create action payload.
type CreateUsersPayload struct {
	// 1 = active, 0 = inactive
	Active string `json:"active" xml:"active"`
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
func (ctx *CreateUsersContext) Created(resp *User) error {
	ctx.Header().Set("Content-Type", "application/vnd.user+json")
	return ctx.Respond(201, resp)
}

// BadRequest sends a HTTP response with status code 400.
func (ctx *CreateUsersContext) BadRequest() error {
	return ctx.RespondBytes(400, nil)
}

// InternalServerError sends a HTTP response with status code 500.
func (ctx *CreateUsersContext) InternalServerError() error {
	return ctx.RespondBytes(500, nil)
}

// DeleteUsersContext provides the users delete action context.
type DeleteUsersContext struct {
	*goa.Context
	UserId string
}

// NewDeleteUsersContext parses the incoming request URL and body, performs validations and creates the
// context used by the users controller delete action.
func NewDeleteUsersContext(c *goa.Context) (*DeleteUsersContext, error) {
	var err error
	ctx := DeleteUsersContext{Context: c}
	rawUserId := c.Get("user_id")
	if rawUserId != "" {
		ctx.UserId = rawUserId
	}
	return &ctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *DeleteUsersContext) OK(resp *User) error {
	ctx.Header().Set("Content-Type", "application/vnd.user+json")
	return ctx.Respond(200, resp)
}

// BadRequest sends a HTTP response with status code 400.
func (ctx *DeleteUsersContext) BadRequest() error {
	return ctx.RespondBytes(400, nil)
}

// InternalServerError sends a HTTP response with status code 500.
func (ctx *DeleteUsersContext) InternalServerError() error {
	return ctx.RespondBytes(500, nil)
}

// ListUsersContext provides the users list action context.
type ListUsersContext struct {
	*goa.Context
}

// NewListUsersContext parses the incoming request URL and body, performs validations and creates the
// context used by the users controller list action.
func NewListUsersContext(c *goa.Context) (*ListUsersContext, error) {
	var err error
	ctx := ListUsersContext{Context: c}
	return &ctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *ListUsersContext) OK(resp UserCollection) error {
	ctx.Header().Set("Content-Type", "application/vnd.user+json; type=collection")
	return ctx.Respond(200, resp)
}

// BadRequest sends a HTTP response with status code 400.
func (ctx *ListUsersContext) BadRequest() error {
	return ctx.RespondBytes(400, nil)
}

// InternalServerError sends a HTTP response with status code 500.
func (ctx *ListUsersContext) InternalServerError() error {
	return ctx.RespondBytes(500, nil)
}

// ShowUsersContext provides the users show action context.
type ShowUsersContext struct {
	*goa.Context
	UserId string
}

// NewShowUsersContext parses the incoming request URL and body, performs validations and creates the
// context used by the users controller show action.
func NewShowUsersContext(c *goa.Context) (*ShowUsersContext, error) {
	var err error
	ctx := ShowUsersContext{Context: c}
	rawUserId := c.Get("user_id")
	if rawUserId != "" {
		ctx.UserId = rawUserId
	}
	return &ctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *ShowUsersContext) OK(resp *User) error {
	ctx.Header().Set("Content-Type", "application/vnd.user+json")
	return ctx.Respond(200, resp)
}

// BadRequest sends a HTTP response with status code 400.
func (ctx *ShowUsersContext) BadRequest() error {
	return ctx.RespondBytes(400, nil)
}

// NotFound sends a HTTP response with status code 404.
func (ctx *ShowUsersContext) NotFound() error {
	return ctx.RespondBytes(404, nil)
}

// InternalServerError sends a HTTP response with status code 500.
func (ctx *ShowUsersContext) InternalServerError() error {
	return ctx.RespondBytes(500, nil)
}

//************************************************************************//
// API "ah_profile": Application Media Types
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

// A user
// Identifier: application/vnd.user+json
type User struct {
	// 1 = active, 0 = inactive
	Active *string `json:"active,omitempty" xml:"active,omitempty"`
	// The date the record was created, in UTC
	CreatedAt *string `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	// The user's email
	Email *string `json:"email,omitempty" xml:"email,omitempty"`
	// The user's first name
	FirstName *string `json:"firstName,omitempty" xml:"firstName,omitempty"`
	// The primary key
	ID *int `json:"id,omitempty" xml:"id,omitempty"`
	// The user's last name
	LastName *string `json:"lastName,omitempty" xml:"lastName,omitempty"`
	// The date the record was updated, in UTC
	ModifiedAt *string `json:"modifiedAt,omitempty" xml:"modifiedAt,omitempty"`
	// The user's phone number
	Phone *string `json:"phone,omitempty" xml:"phone,omitempty"`
	// The user's prescription
	Prescription *string `json:"prescription,omitempty" xml:"prescription,omitempty"`
}

// Validate validates the media type instance.
func (mt *User) Validate() (err error) {
	if mt.Active != nil {
		if ok := goa.ValidatePattern(`^(0|1)$`, *mt.Active); !ok {
			err = goa.InvalidPatternError(`response.active`, *mt.Active, `^(0|1)$`, err)
		}
	}
	if mt.CreatedAt != nil {
		if err2 := goa.ValidateFormat(goa.FormatDateTime, *mt.CreatedAt); err2 != nil {
			err = goa.InvalidFormatError(`response.createdAt`, *mt.CreatedAt, goa.FormatDateTime, err2, err)
		}
	}
	if mt.Email != nil {
		if err2 := goa.ValidateFormat(goa.FormatEmail, *mt.Email); err2 != nil {
			err = goa.InvalidFormatError(`response.email`, *mt.Email, goa.FormatEmail, err2, err)
		}
	}
	if mt.ModifiedAt != nil {
		if err2 := goa.ValidateFormat(goa.FormatDateTime, *mt.ModifiedAt); err2 != nil {
			err = goa.InvalidFormatError(`response.modifiedAt`, *mt.ModifiedAt, goa.FormatDateTime, err2, err)
		}
	}
	if mt.Phone != nil {
		if ok := goa.ValidatePattern(`^\d{10,}$`, *mt.Phone); !ok {
			err = goa.InvalidPatternError(`response.phone`, *mt.Phone, `^\d{10,}$`, err)
		}
	}
	return
}

// , default view
// Identifier: application/vnd.user+json; type=collection
type UserCollection []*User

// Validate validates the media type instance.
func (mt UserCollection) Validate() (err error) {
	for _, e := range mt {
		if e.Active != nil {
			if ok := goa.ValidatePattern(`^(0|1)$`, *e.Active); !ok {
				err = goa.InvalidPatternError(`response[*].active`, *e.Active, `^(0|1)$`, err)
			}
		}
		if e.CreatedAt != nil {
			if err2 := goa.ValidateFormat(goa.FormatDateTime, *e.CreatedAt); err2 != nil {
				err = goa.InvalidFormatError(`response[*].createdAt`, *e.CreatedAt, goa.FormatDateTime, err2, err)
			}
		}
		if e.Email != nil {
			if err2 := goa.ValidateFormat(goa.FormatEmail, *e.Email); err2 != nil {
				err = goa.InvalidFormatError(`response[*].email`, *e.Email, goa.FormatEmail, err2, err)
			}
		}
		if e.ModifiedAt != nil {
			if err2 := goa.ValidateFormat(goa.FormatDateTime, *e.ModifiedAt); err2 != nil {
				err = goa.InvalidFormatError(`response[*].modifiedAt`, *e.ModifiedAt, goa.FormatDateTime, err2, err)
			}
		}
		if e.Phone != nil {
			if ok := goa.ValidatePattern(`^\d{10,}$`, *e.Phone); !ok {
				err = goa.InvalidPatternError(`response[*].phone`, *e.Phone, `^\d{10,}$`, err)
			}
		}
	}
	return
}

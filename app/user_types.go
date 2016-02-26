//************************************************************************//
// API "ah_profile": Application User Types
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

// UserPayload type
type UserPayload struct {
	// 1 = active, 0 = inactive
	Active *string `json:"active,omitempty" xml:"active,omitempty"`
	// The date the record was created, in UTC
	CreatedAt *string `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	// The user's email
	Email *string `json:"email,omitempty" xml:"email,omitempty"`
	// The user's first name
	FirstName *string `json:"firstName,omitempty" xml:"firstName,omitempty"`
	// The uuid
	ID *string `json:"id,omitempty" xml:"id,omitempty"`
	// The user's last name
	LastName *string `json:"lastName,omitempty" xml:"lastName,omitempty"`
	// The date the record was updated, in UTC
	ModifiedAt *string `json:"modifiedAt,omitempty" xml:"modifiedAt,omitempty"`
	// The user's phone number
	Phone *string `json:"phone,omitempty" xml:"phone,omitempty"`
	// The user's prescription
	Prescription *string `json:"prescription,omitempty" xml:"prescription,omitempty"`
}

// Validate validates the type instance.
func (ut *UserPayload) Validate() (err error) {
	if ut.Active != nil {
		if ok := goa.ValidatePattern(`^(0|1)$`, *ut.Active); !ok {
			err = goa.InvalidPatternError(`response.active`, *ut.Active, `^(0|1)$`, err)
		}
	}
	if ut.CreatedAt != nil {
		if err2 := goa.ValidateFormat(goa.FormatDateTime, *ut.CreatedAt); err2 != nil {
			err = goa.InvalidFormatError(`response.createdAt`, *ut.CreatedAt, goa.FormatDateTime, err2, err)
		}
	}
	if ut.Email != nil {
		if err2 := goa.ValidateFormat(goa.FormatEmail, *ut.Email); err2 != nil {
			err = goa.InvalidFormatError(`response.email`, *ut.Email, goa.FormatEmail, err2, err)
		}
	}
	if ut.ModifiedAt != nil {
		if err2 := goa.ValidateFormat(goa.FormatDateTime, *ut.ModifiedAt); err2 != nil {
			err = goa.InvalidFormatError(`response.modifiedAt`, *ut.ModifiedAt, goa.FormatDateTime, err2, err)
		}
	}
	if ut.Phone != nil {
		if ok := goa.ValidatePattern(`^\d{10,}$`, *ut.Phone); !ok {
			err = goa.InvalidPatternError(`response.phone`, *ut.Phone, `^\d{10,}$`, err)
		}
	}
	return
}

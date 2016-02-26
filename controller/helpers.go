package controller

import (
	"time"

	"github.com/alliancehealth/ah_profile/app"
	"github.com/alliancehealth/ah_profile/db"
)

func (c *UsersController) toUserMedia(m db.UserModel) *app.User {
	ca := c.timeToString(&m.CreatedAt)
	ma := c.timeToString(&m.ModifiedAt)
	u := app.User{
		ID:           &m.ID,
		Active:       &m.Active,
		Email:        &m.Email,
		FirstName:    &m.FirstName,
		LastName:     &m.LastName,
		Phone:        &m.Phone,
		Prescription: &m.Prescription,
		CreatedAt:    &ca,
		ModifiedAt:   &ma,
	}

	return &u
}

func (c *UsersController) toUserModel(p *app.UserPayload) (*db.UserModel, error) {
	ca, err := c.stringToTime(p.CreatedAt)
	if err != nil {
		return nil, err
	}
	ma, err := c.stringToTime(p.ModifiedAt)
	if err != nil {
		return nil, err
	}
	u := db.UserModel{
		ID:           *p.ID,
		Active:       *p.Active,
		Email:        *p.Email,
		FirstName:    *p.FirstName,
		LastName:     *p.LastName,
		Phone:        *p.Phone,
		Prescription: *p.Prescription,
		CreatedAt:    ca,
		ModifiedAt:   ma,
	}

	return &u, nil
}

func (c *UsersController) stringToTime(s *string) (time.Time, error) {
	str := *s
	t, err := time.Parse(time.RFC3339, str)
	if err != nil {
		return t, err
	}
	return t, nil
}

func (c *UsersController) timeToString(t *time.Time) string {
	ts := t.Format(time.RFC3339)
	return ts
}

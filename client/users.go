package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// CreateUsersPayload is the data structure used to initialize the users create request body.
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

// Create a new user
func (c *Client) CreateUsers(path string, payload *CreateUsersPayload) (*http.Response, error) {
	var body io.Reader
	b, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("failed to serialize body: %s", err)
	}
	body = bytes.NewBuffer(b)
	u := url.URL{Host: c.Host, Scheme: c.Scheme, Path: path}
	req, err := http.NewRequest("POST", u.String(), body)
	if err != nil {
		return nil, err
	}
	header := req.Header
	header.Set("Content-Type", "application/json")
	return c.Client.Do(req)
}

// Delete a user
func (c *Client) DeleteUsers(path string) (*http.Response, error) {
	var body io.Reader
	u := url.URL{Host: c.Host, Scheme: c.Scheme, Path: path}
	req, err := http.NewRequest("DELETE", u.String(), body)
	if err != nil {
		return nil, err
	}
	header := req.Header
	header.Set("Content-Type", "application/json")
	return c.Client.Do(req)
}

// Show all users
func (c *Client) ListUsers(path string) (*http.Response, error) {
	var body io.Reader
	u := url.URL{Host: c.Host, Scheme: c.Scheme, Path: path}
	req, err := http.NewRequest("GET", u.String(), body)
	if err != nil {
		return nil, err
	}
	header := req.Header
	header.Set("Content-Type", "application/json")
	return c.Client.Do(req)
}

// Show a user
func (c *Client) ShowUsers(path string) (*http.Response, error) {
	var body io.Reader
	u := url.URL{Host: c.Host, Scheme: c.Scheme, Path: path}
	req, err := http.NewRequest("GET", u.String(), body)
	if err != nil {
		return nil, err
	}
	header := req.Header
	header.Set("Content-Type", "application/json")
	return c.Client.Do(req)
}

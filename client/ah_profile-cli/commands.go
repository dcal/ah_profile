package main

import (
	"encoding/json"
	"fmt"
	"github.com/alliancehealth/ah_profile/client"
	"github.com/spf13/cobra"
)

type (
	// CreateUsersCommand is the command line data structure for the create action of users
	CreateUsersCommand struct {
		Payload string
	}

	// DeleteUsersCommand is the command line data structure for the delete action of users
	DeleteUsersCommand struct {
	}

	// ListUsersCommand is the command line data structure for the list action of users
	ListUsersCommand struct {
	}

	// ShowUsersCommand is the command line data structure for the show action of users
	ShowUsersCommand struct {
	}
)

// Run makes the HTTP request corresponding to the CreateUsersCommand command.
func (cmd *CreateUsersCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = "/users"
	}
	var payload client.CreateUsersPayload
	if cmd.Payload != "" {
		err := json.Unmarshal([]byte(cmd.Payload), &payload)
		if err != nil {
			return fmt.Errorf("failed to deserialize payload: %s", err)
		}
	}
	resp, err := c.CreateUsers(path, &payload)
	if err != nil {
		return err
	}
	HandleResponse(c, resp)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *CreateUsersCommand) RegisterFlags(cc *cobra.Command) {
	cc.Flags().StringVar(&cmd.Payload, "payload", "", "Request JSON body")
}

// Run makes the HTTP request corresponding to the DeleteUsersCommand command.
func (cmd *DeleteUsersCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		return fmt.Errorf("missing path argument")
	}
	resp, err := c.DeleteUsers(path)
	if err != nil {
		return err
	}
	HandleResponse(c, resp)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *DeleteUsersCommand) RegisterFlags(cc *cobra.Command) {
}

// Run makes the HTTP request corresponding to the ListUsersCommand command.
func (cmd *ListUsersCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = "/users"
	}
	resp, err := c.ListUsers(path)
	if err != nil {
		return err
	}
	HandleResponse(c, resp)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *ListUsersCommand) RegisterFlags(cc *cobra.Command) {
}

// Run makes the HTTP request corresponding to the ShowUsersCommand command.
func (cmd *ShowUsersCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		return fmt.Errorf("missing path argument")
	}
	resp, err := c.ShowUsers(path)
	if err != nil {
		return err
	}
	HandleResponse(c, resp)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *ShowUsersCommand) RegisterFlags(cc *cobra.Command) {
}

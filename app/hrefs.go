//************************************************************************//
// API "ah_profile": Application Resource Href Factories
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

import "fmt"

// UsersHref returns the resource href.
func UsersHref(user_id interface{}) string {
	return fmt.Sprintf("/users/%v", user_id)
}

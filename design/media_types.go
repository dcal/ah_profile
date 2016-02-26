package design

import (
	. "github.com/goadesign/goa/design/apidsl"
)

var UserMedia = MediaType("application/vnd.user+json", func() {
	Description("A user")
	Reference(UserPayload)
	Attributes(func() {
		Attribute("id")
		Attribute("firstName")
		Attribute("lastName")
		Attribute("phone")
		Attribute("email")
		Attribute("prescription")
		Attribute("active")
		Attribute("createdAt")
		Attribute("modifiedAt")
	})
	View("default", func() {
		Attribute("id")
		Attribute("firstName")
		Attribute("lastName")
		Attribute("phone")
		Attribute("email")
		Attribute("prescription")
		Attribute("active")
		Attribute("createdAt")
		Attribute("modifiedAt")
	})
})

package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var UserPayload = Type("UserPayload", func() {
	Attribute("id", Integer, "The primary key")
	Attribute("firstName", String, "The user's first name")
	Attribute("lastName", String, "The user's last name")
	Attribute("prescription", String, "The user's prescription")
	Attribute("active", String, "1 = active, 0 = inactive", func() {
		Pattern("^(0|1)$")
	})
	Attribute("phone", String, "The user's phone number", func() {
		Pattern("^\\d{10,}$")
	})
	Attribute("email", String, "The user's email", func() {
		Format("email")
	})
	Attribute("createdAt", String, "The date the record was created, in UTC", func() {
		Format("date-time")
	})
	Attribute("modifiedAt", String, "The date the record was updated, in UTC", func() {
		Format("date-time")
	})
})

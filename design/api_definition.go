package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = API("ah_profile", func() {
	Title("Sample Profile Service")
	Description("What diseases do you have or wish you had?")
	Host("localhost:8080")
	Consumes("application/json")
	Consumes("application/xml")
	Produces("application/json")
	Produces("application/xml")
	Scheme("http")

	ResponseTemplate(OK, func(mt string) {
		Description("Resource found.")
		Status(200)
		Media(mt)
	})

	ResponseTemplate(Created, func(mt string) {
		Description("Resource created.")
		Status(201)
		Media(mt)
	})

	ResponseTemplate(InternalServerError, func(mt string) {
		Description("An internal error occurred")
		Status(500)
		Media(mt)
	})
})

package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = Resource("users", func() {
	DefaultMedia(UserMedia)
	BasePath("/users")
	Action("create", func() {
		Routing(
			POST(""),
		)
		Description("Create a new user")
		Payload(func() {
			Member("firstName")
			Member("lastName")
			Member("phone")
			Member("email")
			Member("prescription")
			Member("active")
			Member("created_at")
			Required("firstName")
			Required("lastName")
			Required("phone")
			Required("active")
		})
		Response(Created, func() {
			Media(UserMedia)
		})
		Response(BadRequest)
		Response(InternalServerError)
	})
	Action("show", func() {
		Routing(
			GET("/:user_id"),
		)
		Description("Show a user")
		Params(func() {
			Param("user_id", Integer)
		})
		Response(OK, func() {
			Media(UserMedia)
		})
		Response(BadRequest)
		Response(NotFound)
		Response(InternalServerError)
	})
	Action("list", func() {
		Routing(
			GET(""),
		)
		Description("Show all users")
		Response(OK, func() {
			Media(
				CollectionOf(
					UserMedia, func() {
						View("default")
					}),
			)
		})
		Response(BadRequest)
		Response(InternalServerError)
	})
	Action("delete", func() {
		Routing(
			DELETE("/:user_id"),
		)
		Description("Delete a user")
		Params(func() {
			Param("user_id", Integer)
		})
		Response(OK, func() {
			Media(UserMedia)
		})
		Response(BadRequest)
		Response(InternalServerError)
	})
})

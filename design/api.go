package design

import (
	. "goa.design/goa/v3/dsl"
)

var _ = API("User", func() {
	Title("User Service")
	Description("Service for getting user info")
	Server("user", func() {
		Host("localhost", func() {
			URI("http://localhost:8000")
		})
	})
})

var _ = Service("user", func() {
	Description("The user service provide user details.")

	Method("get", func() {
		Description("Returns User details")
		Result(User)

		HTTP(func() {
			GET("/user")
			Response(StatusOK)
		})
	})

	Files("/openapi.json", "./gen/http/openapi.json")
})

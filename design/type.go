package design

import (
	. "goa.design/goa/v3/dsl"
)

var Company = ResultType("application/vnd.user.company", func() {
	TypeName("Company")

	Attribute("id", UInt, "ID is the unique id of the company")
	Attribute("name", String, "Name of company")
	Attribute("location", String, "Location of company")

	View("tiny", func() {
		Attribute("name")
	})

	View("default", func() {
		Attribute("id")
		Attribute("name")
		Attribute("location")
	})

	Required("id", "name", "location")
})

var User = ResultType("application/vnd.user.user", func() {
	TypeName("User")

	Attribute("id", UInt, "ID is the unique id of the user")
	Attribute("name", String, "Name of user")
	Attribute("latestCompany", Company, "latest company of user")
	Attribute("companies", ArrayOf(Company), "all companies user worked at")

	View("default", func() {
		Attribute("id")
		Attribute("name")
		Attribute("latestCompany", func() {
			View("tiny")
		})
		Attribute("companies", func() {
			View("default")
		})
	})

	Required("id", "name", "latestCompany", "companies")
})

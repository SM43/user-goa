// Code generated by goa v3.2.0, DO NOT EDIT.
//
// user HTTP server types
//
// Command:
// $ goa gen github.com/sm43/user-goa/design

package server

import (
	userviews "github.com/sm43/user-goa/gen/user/views"
)

// GetResponseBody is the type of the "user" service "get" endpoint HTTP
// response body.
type GetResponseBody struct {
	// ID is the unique id of the user
	ID uint `form:"id" json:"id" xml:"id"`
	// Name of user
	Name string `form:"name" json:"name" xml:"name"`
	// latest company of user
	LatestCompany *CompanyResponseBodyTiny `form:"latestCompany" json:"latestCompany" xml:"latestCompany"`
	// all companies user worked at
	Companies []*CompanyResponseBodyTiny `form:"companies" json:"companies" xml:"companies"`
}

// CompanyResponseBodyTiny is used to define fields on response body types.
type CompanyResponseBodyTiny struct {
	// Name of company
	Name string `form:"name" json:"name" xml:"name"`
}

// NewGetResponseBody builds the HTTP response body from the result of the
// "get" endpoint of the "user" service.
func NewGetResponseBody(res *userviews.UserView) *GetResponseBody {
	body := &GetResponseBody{
		ID:   *res.ID,
		Name: *res.Name,
	}
	if res.LatestCompany != nil {
		body.LatestCompany = marshalUserviewsCompanyViewToCompanyResponseBodyTiny(res.LatestCompany)
	}
	if res.Companies != nil {
		body.Companies = make([]*CompanyResponseBodyTiny, len(res.Companies))
		for i, val := range res.Companies {
			body.Companies[i] = marshalUserviewsCompanyViewToCompanyResponseBodyTiny(val)
		}
	}
	return body
}
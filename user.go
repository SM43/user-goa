package userapi

import (
	"context"
	"log"

	user "github.com/sm43/user-goa/gen/user"
)

// user service example implementation.
// The example methods log the requests and return zero values.
type usersrvc struct {
	logger *log.Logger
}

// NewUser returns the user service implementation.
func NewUser(logger *log.Logger) user.Service {
	return &usersrvc{logger}
}

// Returns User details
func (s *usersrvc) Get(ctx context.Context) (res *user.User, err error) {
	s.logger.Print("user.get")

	var companies []*user.Company
	companies = append(companies, &user.Company{ID: 1, Name: "Symantec", Location: "India"})
	companies = append(companies, &user.Company{ID: 2, Name: "Red Hat", Location: "India"})

	// <------Case 1 works------->

	// res = &user.User{
	// 	ID:   1,
	// 	Name: "abc",
	// 	LatestCompany: &user.Company{
	// 		Name: "Redhat",
	// 	},
	// 	Companies: companies,
	// }

	// <------Case 2 works------->

	// res = &user.User{
	// 	ID:   1,
	// 	Name: "abc",
	// 	LatestCompany: &user.Company{
	// 		Name: "Redhat",
	// 	},
	// 	// Companies: companies,
	// }

	// <------Case 3 works------->

	res = &user.User{
		ID:   1,
		Name: "abc",
		// LatestCompany: &user.Company{
		// 	Name: "Redhat",
		// },
		Companies: companies,
	}

	return res, nil
}

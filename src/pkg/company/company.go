package company

import (
	"api.example.com/pkg/user"
	"time"
)

// Company ID
type ID int

// Company Owner-ID
type OwnerID = user.ID

// Company Name
type Name string

func (n Name) valid() bool {
	l := len(n)
	return l > 0 && l < 256
}

type Company struct {
	ID        ID
	Name      Name
	OwnerID   OwnerID
	UpdatedAt time.Time
}

func New(name Name, ownerid OwnerID) *Company {
	return &Company{
		Name:    name,
		OwnerID: ownerid,
	}
}

func (c *Company) validCreate() bool {
	return c.Name.valid() && c.OwnerID.Valid()
}

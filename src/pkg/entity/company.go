package entity

import (
	"fmt"
)

type CompanyID int

// Company Entity
type Company struct {
	ID   CompanyID
	Name string
}

func NewCompany(name string) *Company {
	return &Company{
		Name: name,
	}
}

func (c *Company) Validate() error {
	if l := len(c.Name); l == 0 {
		return fmt.Errorf("entity.Company.Validate: Name.Length %d ≥ 1", l)
	}
	return nil
}

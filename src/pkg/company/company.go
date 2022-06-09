package company

type ID int

func (id ID) valid() bool {
	return id > 0
}

type Name string

// Name ≥ 1
func (n Name) valid() bool {
	return len(n) >= 1
}

type Company struct {
	ID   ID
	Name Name
}

func New(n Name) *Company {
	return &Company{
		Name: n,
	}
}

func (c *Company) valid() bool {
	return c.Name.valid()
}

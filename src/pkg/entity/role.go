package entity

type RoleID int

type Role struct {
	Company *Company
	ID      RoleID
	Name    string
}

func NewRoleAdmin(company *Company) *Role {
	return &Role{
		Company: company,
		Name:    "管理者",
	}
}

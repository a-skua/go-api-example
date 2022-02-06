package repository

import (
	"api.example.com/pkg/entity"
)

// テスト用 Mock
type Mock struct {
	User    *entity.User
	Company *entity.Company
	Role    *entity.Role
	Error   error
	Tx      Tx
}

func (r *Mock) Transaction() Tx {
	return r.Tx
}

func (r *Mock) UserCreate(*entity.User) (*entity.User, error) {
	return r.User, r.Error
}

func (r *Mock) UserRead(entity.UserID) (*entity.User, error) {
	return r.User, r.Error
}

func (r *Mock) UserUpdate(*entity.User) (*entity.User, error) {
	return r.User, r.Error
}

func (r *Mock) UserDelete(entity.UserID) error {
	return r.Error
}

func (r *Mock) CompanyCreateTx(*entity.Company, Tx) (*entity.Company, Tx) {
	return r.Company, r.Tx
}

func (r *Mock) CompanyAddEmployeeTx(entity.CompanyID, entity.UserID, Tx) Tx {
	return r.Tx
}

func (r *Mock) RoleCreateTx(*entity.Role, Tx) (*entity.Role, Tx) {
	return r.Role, r.Tx
}

func (r *Mock) EmployeeAddRoleTx(entity.CompanyID, entity.UserID, entity.RoleID, Tx) Tx {
	return r.Tx
}

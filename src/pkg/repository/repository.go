package repository

import (
	"api.example.com/pkg/entity"
)

type Repository interface {
	transaction
	User
	Company
}

// transaction を管理
//
//     var tx Transaction
//     err := tx.Error()
//     if err != nil {
//         tx.Rollback()
//         return
//     }
//     tx.Commit()
type Tx interface {
	Rollback() error
	Commit() error
	Error() error
}

type transaction interface {
	Transaction() Tx
}

type User interface {
	UserCreate(*entity.User) (*entity.User, error)
	UserRead(entity.UserID) (*entity.User, error)
	UserUpdate(*entity.User) (*entity.User, error)
	UserDelete(entity.UserID) error
}

type Company interface {
	transaction
	CompanyCreateTx(*entity.Company, Tx) (*entity.Company, Tx)
	CompanyAddEmployeeTx(entity.CompanyID, entity.UserID, Tx) Tx
	RoleCreateTx(*entity.Role, Tx) (*entity.Role, Tx)
	EmployeeAddRoleTx(entity.CompanyID, entity.UserID, entity.RoleID, Tx) Tx
}

package user

type Repository interface {
	UserCreate(*User) (*User, error)
	UserRead(ID) (*User, error)
	UserUpdate(*User) (*User, error)
	UserDelete(ID) error
}

package company

import (
	"errors"
)

type Repository interface {
	CompanyCreate(*Company) (*Company, error)
}

type Server interface {
	Create(*Company) (*Company, error)
}

// impl Server
type server struct {
	repository Repository
}

func NewServer(repo Repository) Server {
	return &server{repo}
}

func (s *server) Create(c *Company) (*Company, error) {
	if ok := c.validCreate(); !ok {
		return nil, errors.New("invalid create")
	}

	return s.repository.CompanyCreate(c)
}

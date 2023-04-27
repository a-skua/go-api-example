package company

import (
	"errors"
	"fmt"
)

type Repository interface {
	CompanyCreate(*Company) (*Company, error)
	CompanyRead(ID) (*Company, error)
}

type Server interface {
	Create(*Company) (*Company, error)
	Read(ID) (*Company, error)
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

func (s *server) Read(id ID) (*Company, error) {
	ok := id.Valid()

	if !ok {
		return nil, fmt.Errorf("pkg/comopany.Read: invalid company_id")
	}

	return s.repository.CompanyRead(id)
}

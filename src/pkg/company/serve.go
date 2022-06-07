package company

import (
	"fmt"
)

type Repository interface {
	CompanyCreate(*Company) (*Company, error)
	CompanyRead(ID) (*Company, error)
	CompanyUpdate(*Company) (*Company, error)
	CompanyDelete(ID) error
}

type Server interface {
	Create(*Company) (*Company, error)
	Read(ID) (*Company, error)
	Update(*Company) (*Company, error)
	Delete(ID) error
}

// impl Server
type server struct {
	repository Repository
}

func NewServer(repo Repository) Server {
	return &server{repo}
}

func (s *server) Create(c *Company) (*Company, error) {
	ok := c.valid()
	if !ok {
		return nil, fmt.Errorf("pkg/company.Create: invalid company")
	}

	c, err := s.repository.CompanyCreate(c)
	if err != nil {
		return nil, fmt.Errorf("pkg/company.Create: %w", err)
	}

	return c, nil
}

func (s *server) Read(id ID) (*Company, error) {
	ok := id.valid()
	if !ok {
		return nil, fmt.Errorf("pkg/company.Read: invalid id")
	}

	c, err := s.repository.CompanyRead(id)
	if err != nil {
		return nil, fmt.Errorf("pkg/company.Read: %w", err)
	}

	return c, nil
}

func (s *server) Update(c *Company) (*Company, error) {
	ok := c.ID.valid() && c.valid()
	if !ok {
		return nil, fmt.Errorf("pkg/company.Update: invalid company")
	}

	c, err := s.repository.CompanyUpdate(c)
	if err != nil {
		return nil, fmt.Errorf("pkg/company.Update: %w", err)
	}

	return c, nil
}

func (s *server) Delete(id ID) error {
	ok := id.valid()
	if !ok {
		return fmt.Errorf("pkg/company.Delete: invalid id")
	}

	err := s.repository.CompanyDelete(id)
	if err != nil {
		return fmt.Errorf("pkg/company.Delete: %w", err)
	}

	return nil
}

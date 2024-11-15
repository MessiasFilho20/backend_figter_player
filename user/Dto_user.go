package user

import "errors"

type DtoUser struct {
	Name   string `json:"Name"`
	Email  string `json:"Email"`
	Credit string `json:"credit"`
}

func (c *DtoUser) UserValidate() error {

	if c.Name == "" {
		return errors.New("campo nome está vazio")
	}

	if c.Email == "" {
		return errors.New("o campo email está vazio")
	}

	if c.Credit == "" {
		return errors.New("campo credit vazio ")
	}

	return nil
}

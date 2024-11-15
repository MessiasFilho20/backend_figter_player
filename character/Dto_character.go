package character

import (
	"errors"
	"strings"
)

type DtoCharacter struct {
	UserID uint   `josn:"userID"`
	Name   string `json:"Name"`
	Sexo   string `josn:"Sexo"`
	Classe string `jaon:"Classe"`
}

func (c *DtoCharacter) ValidCharacter() error {
	if c.Name == "" {
		return errors.New("campo nome Vazio")
	}

	sexoValid := map[string]bool{"masculino": true, "feminino": true}

	if strings.TrimSpace(c.Sexo) == "" {
		return errors.New("campo sexo Vazio")
	} else if !sexoValid[strings.ToLower(c.Sexo)] {
		return errors.New("campo 'Sexo' inválido: use 'masculino' ou 'feminino'")
	}

	if c.Classe == "" {
		return errors.New("campo classe Vazio")
	}

	if c.UserID == 0 {
		return errors.New("campo 'UserID' inválido: deve ser maior que 0")
	}

	return nil
}

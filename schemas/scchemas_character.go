package schemas

type Character struct {
	ID     uint   `gorm:"primaryKey"`
	UserID uint   `josn:"userID"`
	Name   string `json:"Name"`
	Sexo   string `josn:"Sexo"`
	Classe string `jaon:"Classe"`
}

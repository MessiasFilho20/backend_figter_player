package schemas

type User struct {
	ID     uint   `gorm:"primaryKey"`
	Name   string `json:"Name"`
	Email  string `json:"Email"`
	Credit string `json:"credit"`
}

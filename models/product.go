package models

type Product struct {
	Id    uint   `gorm:"primaryKey autoIncrement" json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}

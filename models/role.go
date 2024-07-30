package models

type Role struct {
	ID    uint   `gorm:"primaryKey"`
	Name  string `json:"name" gorm:"unique;not null"`
	Pages []Page `gorm:"many2many:role_pages"`
}

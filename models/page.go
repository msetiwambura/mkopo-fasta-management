package models

type Page struct {
	ID          uint   `gorm:"primaryKey"`
	Name        string `json:"name"`
	URL         string `json:"url" gorm:"unique;not null"`
	Description string `json:"description"`
}

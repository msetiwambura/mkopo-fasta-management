package models

import "time"

type User struct {
	ID        uint      `gorm:"primaryKey"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Email     string    `json:"email" gorm:"unique;not null"`
	Password  string    `json:"password"`
	RoleID    uint      `json:"roleId"`
	Role      Role      `gorm:"foreignKey:RoleID"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

type ResUser struct {
	ID        uint   `gorm:"primaryKey"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email" gorm:"unique;not null"`
	RoleID    uint   `json:"roleId"`
}
type Login struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

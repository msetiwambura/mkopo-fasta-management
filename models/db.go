package models

import "usrmanagement/configs"

func MigrateDB() {
	err := configs.DB.AutoMigrate(
		&User{},
		&Role{},
	)
	if err != nil {
		return
	}
}

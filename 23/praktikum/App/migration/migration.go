package migration

import (
	"20/model"

	"gorm.io/gorm"
)

func InitMigrationMysql(db *gorm.DB) {
	db.AutoMigrate(&model.User{})
	// auto migrate untuk table book
}

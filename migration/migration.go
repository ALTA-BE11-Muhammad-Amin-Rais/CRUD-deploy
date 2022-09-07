package migration

import (
	//_mAuth "Belajar/CleanCode/features/authorized/data"
	_mBooks "Belajar/CleanCode/features/book/data"
	_mUsers "Belajar/CleanCode/features/user/data"

	"gorm.io/gorm"
)

func InitMigrate(db *gorm.DB) {

	db.AutoMigrate(&_mUsers.User{})
	db.AutoMigrate(&_mBooks.Book{})

}

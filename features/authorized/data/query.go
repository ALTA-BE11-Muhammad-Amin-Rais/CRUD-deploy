package data

import (
	"Belajar/CleanCode/features/authorized"

	"log"

	"gorm.io/gorm"
)

type userLogin struct {
	DB *gorm.DB
}

func New(conn *gorm.DB) authorized.DataInterface {

	return &userLogin{
		DB: conn,
	}

}

func (repo *userLogin) LoginUser(email, password string) (authorized.AuthCore, error) {

	var auth User
	txEmail := repo.DB.Where("email = ?", email).First(&auth)
	if txEmail.Error != nil {
		log.Println("Error tx")
		return authorized.AuthCore{}, txEmail.Error
	}

	if txEmail.RowsAffected != 1 {
		log.Println("Error row")
		return authorized.AuthCore{}, txEmail.Error
	}

	var data = toCore(auth)

	return data, nil

}

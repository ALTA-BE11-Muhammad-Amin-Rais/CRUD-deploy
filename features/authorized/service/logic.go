package service

import (
	"Belajar/CleanCode/features/authorized"
	"Belajar/CleanCode/middlewares"
	"log"

	"golang.org/x/crypto/bcrypt"
)

type authorizedUsecase struct {
	authorizedData authorized.DataInterface
}

func New(data authorized.DataInterface) authorized.UsecaseInterface {
	return &authorizedUsecase{
		authorizedData: data,
	}
}

func (usecase *authorizedUsecase) LoginAuthorized(email, password string) (string, error) {

	var err error
	if email == "" || password == "" {
		return "", err
	}

	results, errEmail := usecase.authorizedData.LoginUser(email, password)
	if errEmail != nil {
		return "", errEmail
	}

	errPw := bcrypt.CompareHashAndPassword([]byte(results.Password), []byte(password))
	if errPw != nil {
		log.Println("Error pw")
		return "", err
	}

	token, errToken := middlewares.CreateToken(int(results.ID))

	if errToken != nil {
		return "", err
	}

	if token == "" {
		return "", err
	}

	return token, nil

}

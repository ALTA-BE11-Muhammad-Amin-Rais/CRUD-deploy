package factory

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"

	authData "Belajar/CleanCode/features/authorized/data"
	authDelivery "Belajar/CleanCode/features/authorized/delivery"
	authService "Belajar/CleanCode/features/authorized/service"

	userData "Belajar/CleanCode/features/user/data"
	userDelivery "Belajar/CleanCode/features/user/delivery"
	userService "Belajar/CleanCode/features/user/service"

	bookData "Belajar/CleanCode/features/book/data"
	bookDelivery "Belajar/CleanCode/features/book/delivery"
	bookService "Belajar/CleanCode/features/book/service"
)

func InitFactory(e *echo.Echo, db *gorm.DB) {

	userDataFactory := userData.New(db)
	userUsecaseFactory := userService.New(userDataFactory)
	userDelivery.New(e, userUsecaseFactory)

	authDataFactory := authData.New(db)
	authUsecaseFactory := authService.New(authDataFactory)
	authDelivery.New(e, authUsecaseFactory)

	bookDataFactory := bookData.New(db)
	bookUsecaseFactory := bookService.New(bookDataFactory)
	bookDelivery.New(e, bookUsecaseFactory)

}

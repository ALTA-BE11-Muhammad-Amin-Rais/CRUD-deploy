package delivery

import (
	"Belajar/CleanCode/features/authorized"
	"Belajar/CleanCode/utils/helper"

	"github.com/labstack/echo/v4"
)

type AuthHandler struct {
	userUsecase authorized.UsecaseInterface
}

func New(e *echo.Echo, usecase authorized.UsecaseInterface) {

	handler := AuthHandler{
		userUsecase: usecase,
	}

	e.POST("/auth", handler.Auth)

}

func (h *AuthHandler) Auth(c echo.Context) error {

	var req Request
	errBind := c.Bind(&req)
	if errBind != nil {
		return c.JSON(400, errBind)
	}

	str, err := h.userUsecase.LoginAuthorized(req.Email, req.Password)
	if err != nil {
		return c.JSON(404, err)
	}
	return c.JSON(200, helper.SuccessDataResponseHelper("Login Success", str))

}

package controller

import (
	"20/dto"
	"20/middleware"
	"20/model"
	"20/repository"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

// UserController defines the user controller interface.
type UserController interface {
	GetAllUsers(c echo.Context) error
	CreateUser(c echo.Context) error
}

type userController struct {
	userRepo repository.UserRepository
}

// NewUserController creates a new UserController instance.
func NewUserController(uRepo repository.UserRepository) UserController {
	return &userController{
		userRepo: uRepo,
	}
}

// GetAllUsers retrieves all users.
func (u *userController) GetAllUsers(c echo.Context) error {
	users, err := u.userRepo.Find()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, echo.Map{
		"data": users,
	})
}

// CreateUser creates a new user.
func (u *userController) CreateUser(c echo.Context) error {
	var user model.User
	err := c.Bind(&user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": err.Error(),
		})
	}

	err = u.userRepo.Create(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": fmt.Sprintf("failed to create user: %v", err),
		})
	}

	token, err := middleware.CreateToken(user.Email)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": fmt.Sprintf("failed to create token: %v", err),
		})
	}

	uRes := dto.DTOUsers{
		Email: user.Email,
		Token: token,
	}
	return c.JSON(http.StatusOK, echo.Map{
		"data": uRes,
	})
}

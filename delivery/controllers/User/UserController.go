package controllers

import (
	"net/http"
	helpers "project-stokku/delivery/helpers/request"
	views "project-stokku/delivery/helpers/view"
	"project-stokku/entity"
	repository "project-stokku/repository/User"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

type userController struct {
	Connect   repository.UserModel
	Validator *validator.Validate
}

func NewUserController(conn repository.UserModel, valid *validator.Validate) *userController {
	return &userController{
		Connect:  conn,
		Validator: valid,
	}
}

func (u *userController) Create(c echo.Context) error {
	var user helpers.UserRequest

	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, views.BadRequest("Invalid argument exception"))
	}

	if err := u.Validator.Struct(user); err != nil {
		return c.JSON(http.StatusBadRequest, views.InputRequest("is Required", "name, email, password"))
	}

	if status := u.Connect.CheckDuplicate(user.Email); status {
		return c.JSON(http.StatusBadRequest, views.InputRequest("must be replace", "Email is already registered"))
	}

	u.Connect.Create(&entity.User{
		Name: user.Name,
		Email: user.Email,
		Password: user.Password,
	})

	return c.JSON(http.StatusCreated, views.SuccessInsert("Create", "User", user))
}

func CreateToken(userId uint) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["userId"] = userId
	claims["expired"] = time.Now().Add(time.Hour * 1).Unix() //Token expires after 1 hour
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte("$p4ssw0rd"))
}

func (u *userController) Login(c echo.Context) error {
	user := helpers.LoginRequest{}

	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, views.BadRequest("Invalid argument exception"))
	}

	if err := u.Validator.Struct(user); err != nil {
		return c.JSON(http.StatusBadRequest, views.InputRequest("is Required", "email, password"))
	}

	status, data := u.Connect.Login(user.Email, user.Password)
	result := views.LoginResponse{Detail: data}

	if status && result.Token == "" {
		token, _ := CreateToken(data.ID)
		result.Token = token
		return c.JSON(http.StatusOK, views.SuccessResponse("Berhasil", "Login", result))
	}

	return c.JSON(http.StatusUnauthorized, views.Unauthorized())
}
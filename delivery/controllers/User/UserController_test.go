package controllers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"project-stokku/entity"
	"strings"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {
	t.Run("Success Insert", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]interface{}{
			"name": "Galang", "email" : "galang@gmail.com", "password" : "password",
		})
		
		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(string(requestBody)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/register")

		userController := NewUserController(&mockUser{}, validator.New())
		userController.Create(context)

		type response struct {
			Code    int
			Message string
			Data    interface{}
		}

		var resp response

		json.Unmarshal([]byte(res.Body.Bytes()), &resp)
		assert.Equal(t, 201, resp.Code)
		assert.Equal(t, "Success Create User", resp.Message)
		assert.Equal(t, map[string]interface {}(map[string]interface {}{"email":"galang@gmail.com", "name":"Galang", "password":"password"}), resp.Data)
	})

	t.Run("Failed Bind", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]interface{}{
			"name": 1, "email" : "galang@gmail.com", "password" : "password",
		})
		
		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(string(requestBody)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/register")

		userController := NewUserController(&mockUser{}, validator.New())
		userController.Create(context)

		type response struct {
			Code    int
			Message string
			Data    interface{}
		}

		var resp response

		json.Unmarshal([]byte(res.Body.Bytes()), &resp)
		assert.Equal(t, 422, resp.Code)
		assert.Equal(t, "Invalid argument exception", resp.Message)
		assert.Nil(t, resp.Data)
	})	

	t.Run("Error Validasi Create", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]interface{}{
			"name": "Galang",
		})
		
		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(string(requestBody)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/register")

		userController := NewUserController(&mockUser{}, validator.New())
		userController.Create(context)

		type response struct {
			Code    int
			Message string
			Data    interface{}
		}

		var resp response

		json.Unmarshal([]byte(res.Body.Bytes()), &resp)
		assert.Equal(t, 400, resp.Code)
		assert.Equal(t, "Field is Required! name, email, password", resp.Message)
		assert.Nil(t, resp.Data)
	})

	t.Run("Error Duplicate", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]interface{}{
			"name": "Galang", "email" : "admin@gmail.com", "password" : "password",
		})

		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(string(requestBody)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/register")

		userController := NewUserController(&mockErrorUser{}, validator.New())
		userController.Create(context)

		type response struct {
			Code    int
			Message string
			Data    interface{}
		}

		var resp response

		json.Unmarshal([]byte(res.Body.Bytes()), &resp)
		assert.Equal(t, 400, resp.Code)
		assert.Equal(t, "Field must be replace! Email is already registered", resp.Message)
		assert.Nil(t, resp.Data)
	})
}

var token string
func TestLogin(t *testing.T) {
	t.Run("Success Login", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]interface{}{
			"email": "dakasakti.id@gmail.com",
			"password": "password",
		})

		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(string(requestBody)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/login")

		controller := NewUserController(&mockUser{}, validator.New())
		controller.Login(context)

		type ResponseStructure struct {
			Code    int
			Message string
			Data    interface{}
		}

		var response ResponseStructure

		json.Unmarshal([]byte(res.Body.Bytes()), &response)
		assert.NotNil(t, response.Data)
		data := response.Data.(map[string]interface{})
		token = data["Token"].(string)
	})

	t.Run("Failed Bind", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]interface{}{
			"email" : "dakasakti.id@gmail.com",
			"password" : 123,
		})
		
		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(string(requestBody)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/login")

		controller := NewUserController(&mockUser{}, validator.New())
		controller.Login(context)

		type response struct {
			Code    int
			Message string
			Data    interface{}
		}

		var resp response

		json.Unmarshal([]byte(res.Body.Bytes()), &resp)
		assert.Equal(t, 422, resp.Code)
		assert.Equal(t, "Invalid argument exception", resp.Message)
		assert.Nil(t, resp.Data)
	})	

	t.Run("Error Validasi Login", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]interface{}{
			"email" : "dakasakti.id@gmail.com",
		})
		
		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(string(requestBody)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/login")

		userController := NewUserController(&mockUser{}, validator.New())
		userController.Login(context)

		type response struct {
			Code    int
			Message string
			Data    interface{}
		}

		var resp response

		json.Unmarshal([]byte(res.Body.Bytes()), &resp)
		assert.Equal(t, 400, resp.Code)
		assert.Equal(t, "Field is Required! email, password", resp.Message)
		assert.Nil(t, resp.Data)
	})

	t.Run("Error Input Login", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]interface{}{
			"email" : "admin@gmail.com", "password" : "password",
		})

		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(string(requestBody)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/login")

		Controller := NewUserController(&mockErrorUser{}, validator.New())
		Controller.Login(context)

		type response struct {
			Code    int
			Message string
			Data    interface{}
		}

		var resp response

		json.Unmarshal([]byte(res.Body.Bytes()), &resp)
		assert.Equal(t, 401, resp.Code)
		assert.Equal(t, "Email atau Password Salah", resp.Message)
		assert.Nil(t, resp.Data)
	})
}

type mockUser struct{}

func (mur *mockUser) Create(user *entity.User) error {
	return nil
}

func (mur *mockUser) CheckDuplicate(email string) bool {
	return false
}

func (mur *mockUser) Login(email, password string) (bool, entity.User) {
	return true, entity.User{}
}

type mockErrorUser struct{}

func (mur *mockErrorUser) Create(user *entity.User) error {
	return nil
}

func (mur *mockErrorUser) CheckDuplicate(email string) bool {
	return true
}

func (mur *mockErrorUser) Login(email, password string) (bool, entity.User) {
	return false, entity.User{}
}
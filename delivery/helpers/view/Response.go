package views

import (
	"fmt"
	"net/http"
	"project-stokku/entity"
)

type LoginResponse struct {
	Detail  entity.User
	Token	string
}

func Unauthorized() map[string]interface{} {
	return map[string]interface{}{
		"code": http.StatusUnauthorized,
		"message" : "Email atau Password Salah",
		"data" : nil,
	}
}

func BadRequest(message interface{}) map[string]interface{} {
	return map[string]interface{}{
		"code": http.StatusUnprocessableEntity,
		"message" : message,
		"data" : nil,
	}
}

func SuccessResponse(status, model string, data interface{}) map[string]interface{} {
	message := fmt.Sprintf("Success %s %s", status, model)
	
	return map[string]interface{}{
		"code":    http.StatusOK,
		"message": message,
		"data": data,
	}
}

func SuccessInsert(status, model string, data interface{}) map[string]interface{} {
	message := fmt.Sprintf("Success %s %s", status, model)
	
	return map[string]interface{}{
		"code":    http.StatusCreated,
		"message": message,
		"data": data,
	}
}

func NotFoundResponse(input interface{}) map[string]interface{} {
	message := fmt.Sprintf("%s not Found", input)

	return map[string]interface{}{
		"code": http.StatusNotFound,
		"message" : message,
		"data" : nil,
	}
}

func InputRequest(status string, input interface{}) map[string]interface{} {
	message := fmt.Sprintf("Field %s! %s", status, input)
	
	return map[string]interface{}{
		"code": http.StatusBadRequest,
		"message" : message,
		"data" : nil,
	}
}
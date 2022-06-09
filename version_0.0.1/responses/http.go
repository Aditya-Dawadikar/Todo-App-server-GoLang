package responses

import (
	"todo_service/version_0.0.1/models"
)

type LoginSuccess struct {
	Status   int    `json:"status"`
	Message  string `json:"message"`
	Username string `json:"username"`
	Userid   string `json:"userid"`
}

type GeneralSuccess struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type LoginError struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type UnknownError struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type FoundUsers struct {
	Status  int               `json:"status"`
	Message string            `json:"message"`
	Users   []models.TodoUser `json:"users`
}

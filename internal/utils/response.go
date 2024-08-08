package utils

type SuccessResponse struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
	Message string      `json:"msg"`
}

type WarningResponse struct {
	Success bool        `json:"success"`
	Warning string      `json:"warning"`
	Data    interface{} `json:"data"`
	Message string      `json:"msg"`
}

type ErrorResponse struct {
	Success bool   `json:"success"`
	Error   string `json:"error"`
	Message string `json:"msg"`
}

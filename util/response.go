package util

import (
	"encoding/json"
	"net/http"
)

type Error struct {
	Status bool `json:"status"`
	Message string `json:"message"`
	ErrorCode int `json:"error_code"`
	Data interface{}
}

type Success struct {
	Status bool `json:"status"`
	Message string `json:"message"`
	Data interface{} `json:"data"`
}

func ErrorResponse(w http.ResponseWriter, statusCode int, errorCode int, message string, ) {
	w.WriteHeader(statusCode)
	  json.NewEncoder(w).Encode(Error{
		Status: false,
		ErrorCode: errorCode,
		Message: message,
	})
}
func SuccessResponse(w http.ResponseWriter, statusCode int, message string,  data interface{}) {
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(Success{Status: true, Message: message, Data: data})
}
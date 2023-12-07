package logger

import (
	"encoding/json"
	"log"
)

func Error(message, err string) {
	param := Log{
		Type:    "error",
		Message: message,
		Error:   err,
	}
	marshal, _ := json.Marshal(&param)
	log.Println(string(marshal))
}

func Warn(message string) {
	param := Log{
		Type:    "warn",
		Message: message,
	}
	marshal, _ := json.Marshal(&param)
	log.Println(string(marshal))
}

func Info(message string) {
	param := Log{
		Type:    "info",
		Message: message,
	}
	marshal, _ := json.Marshal(&param)
	log.Println(string(marshal))
}

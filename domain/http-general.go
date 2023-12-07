package domain

import (
	"encoding/json"
	"log"
	"net/http"
	"runtime/debug"
	"strings"
)

type GeneralHandlerType func(w http.ResponseWriter, r *http.Request) (GeneralResponse, error)

type GeneralResponse struct {
	Code         int         `json:"code,omitempty"`
	Message      string      `json:"message,omitempty"`
	DeveloperMsg []string    `json:"developer_msg,omitempty"`
	Data         interface{} `json:"data"`
}

func parseMethod(h http.HandlerFunc, method string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if strings.ToLower(r.Method) != method {
			w.WriteHeader(400)
			return
		}
		h(w, r)
	}
}

func GET(h GeneralHandlerType) http.HandlerFunc {
	return parseMethod(HandleFunc(h), "get")
}

func POST(h GeneralHandlerType) http.HandlerFunc {
	return parseMethod(HandleFunc(h), "post")
}

func HandleFunc(h GeneralHandlerType) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if p := recover(); p != nil {
				msg, _ := p.(string)
				log.Println("panic encountered; stacktrace\n", string(debug.Stack()))
				w.Write([]byte(msg))
				http.Error(w, "", http.StatusInternalServerError)
			}
		}()

		resp, err := h(w, r)
		if err != nil {
			resp.Code = http.StatusInternalServerError
			resp.Message = "internal error"
			resp.DeveloperMsg = append(resp.DeveloperMsg, err.Error())
		}
		bResp, _ := json.Marshal(&resp)
		if resp.Code == 0 {
			resp.Code = http.StatusOK
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(resp.Code)
		w.Write(bResp)
	}
}

func ValidateJSONContentType(r *http.Request) bool {
	if r == nil {
		return false
	}

	contentType := strings.ToLower(r.Header.Get("Content-type"))
	return contentType == "application/json"
}

package domain

import (
	"encoding/json"
	"net/http"
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
				w.WriteHeader(500)
				msg, _ := p.(string)
				http.Error(w, "", 500)
				w.Write([]byte(msg))
			}
		}()

		resp, err := h(w, r)
		if err != nil {
			resp.Code = 500
			resp.Message = "internal error"
			resp.DeveloperMsg = append(resp.DeveloperMsg, err.Error())
			http.Error(w, "", resp.Code)
		}
		bResp, _ := json.Marshal(&resp)
		if resp.Code == 0 {
			resp.Code = 200
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(resp.Code)
		w.Write(bResp)
	}
}

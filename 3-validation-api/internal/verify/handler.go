package verify

import (
	"fmt"
	"net/http"
)

type VerifyHandler struct{}

func NewVerifyHandler(router *http.ServeMux) {
	handler := VerifyHandler{}
	router.HandleFunc("POST /send", handler.Send())
	router.HandleFunc("GET /verify/{hash}", handler.Verify())
}

func (handler *VerifyHandler) Send() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		fmt.Println("Send")
	}
}
func (handler *VerifyHandler) Verify() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		fmt.Println("Verify")
	}
}

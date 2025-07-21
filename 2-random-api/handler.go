package main

import (
	"fmt"
	"math/rand"
	"net/http"
)

type RandomNumberHandler struct{}

func NewRandomNumberHandler(router *http.ServeMux) {
	handler := &RandomNumberHandler{}
	router.HandleFunc("/random", handler.CreateRandomNum())
}

func (hanler *RandomNumberHandler) CreateRandomNum() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte(fmt.Sprint(rand.Intn(7))))
	}
}

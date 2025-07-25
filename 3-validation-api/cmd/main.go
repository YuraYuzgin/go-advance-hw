package main

import (
	"emailVerify/3-validation-api/internal/verify"
	"fmt"
	"net/http"
)

func main() {
	router := http.NewServeMux()
	verify.NewVerifyHandler(router)

	server := http.Server{
		Addr:    ":8081",
		Handler: router,
	}

	fmt.Println("Serve is listening on port 8081")
	server.ListenAndServe()
}

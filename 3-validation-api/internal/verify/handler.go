package verify

import (
	"crypto/rand"
	"emailVerify/3-validation-api/pkg/req"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/smtp"
	"os"

	"github.com/jordan-wright/email"
)

type VerifyHandler struct{}

func NewVerifyHandler(router *http.ServeMux) {
	handler := VerifyHandler{}
	router.HandleFunc("POST /send", handler.Send())
	router.HandleFunc("GET /verify/{hash}", handler.Verify())
}

func (handler *VerifyHandler) Send() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Send")
		payload, err := req.HandleBody[Request](&w, r)
		if err != nil {
			return
		}

		b := make([]byte, 16)
		rand.Read(b)
		hash := hex.EncodeToString(b)

		saveHush := Hash{
			Email: payload.Email,
			Hash:  hash,
		}
		// Преобразуем структуру в JSON
		jsonData, err := json.MarshalIndent(saveHush, "", "  ")
		if err != nil {
			panic(err)
		}

		// Сохраняем JSON в файл
		err = os.WriteFile("hash.json", jsonData, 0644)
		if err != nil {
			panic(err)
		}

		e := email.NewEmail()
		e.From = "go-hw@example.com"
		e.To = []string{payload.Email}
		e.Subject = "Email Verification"
		e.Text = []byte("Please click the following link to verify your email: http://localhost:8081/verify/" + hash)
		e.HTML = []byte("<p>Please click the following link to verify your email: <a href=\"http://localhost:8081/verify/" + hash + "\">Verify Email</a></p>")
		err = e.Send("smtp.gmail.com:587", smtp.PlainAuth("", "yurayuzgin@gmail.com", "wizd hpeb grcw ussx", "smtp.gmail.com"))
		if err != nil {
			log.Printf("Failed to send email: %v", err)
			http.Error(w, "Failed to send verification email", http.StatusInternalServerError)
			return
		}
	}
}
func (handler *VerifyHandler) Verify() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Verify")
		hash := r.PathValue("hash")

		fileContent, err := os.ReadFile("hash.json")
		if err != nil {
			panic(err)
		}
		var referenceHash Hash
		err = json.Unmarshal(fileContent, &referenceHash)
		if err != nil {
			panic(err)
		}
		fmt.Println(referenceHash)
		if referenceHash.Hash != hash {
			os.Remove("hash.go")
		}
	}
}

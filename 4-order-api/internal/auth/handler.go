package auth

import (
	"cart-api/configs"
	"cart-api/pkg/jwt"
	"cart-api/pkg/req"
	"cart-api/pkg/res"
	"net/http"
)

type AuthHandlerDeps struct {
	*configs.Config
	*AuthService
}

type AuthHandler struct {
	*configs.Config
	*AuthService
}

func NewAuthHandler(router *http.ServeMux, deps AuthHandlerDeps) {
	handler := &AuthHandler{
		Config:      deps.Config,
		AuthService: deps.AuthService,
	}
	router.HandleFunc("POST /auth/phone", handler.PhoneVerification())
	router.HandleFunc("POST /auth/code", handler.CodeVerification())
}

func (handler *AuthHandler) PhoneVerification() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := req.HandleBody[AuthRequest](&w, r)
		if err != nil {
			return
		}
		user, err := handler.AuthService.AuthByPhone(body.Phone)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}
		data := AuthResponse{
			SessionId: user.SessionId,
		}
		res.Json(w, data, 200)
	}
}

func (handler *AuthHandler) CodeVerification() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := req.HandleBody[CodeVerificationRequest](&w, r)
		if err != nil {
			return
		}
		phone, err := handler.AuthService.CodeVerification(body.SessionId, body.Code)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}
		token, err := jwt.NewJWT(handler.Config.Auth.Secret).Create(phone)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		data := CodeVerificationResponse{
			Token: token,
		}
		res.Json(w, data, 200)
	}
}

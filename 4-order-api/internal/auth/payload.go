package auth

type AuthRequest struct {
	Phone string `json:"phone" validate:"required"`
}

type AuthResponse struct {
	SessionId string `json:"sessionId"`
}

type CodeVerificationRequest struct {
	SessionId string `json:"sessionId" validate:"required"`
	Code      int    `json:"code" validate:"required"`
}

type CodeVerificationResponse struct {
	Token string `json:"token"`
}

package verify

type Request struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password"`
	Address  string `json:"address"`
}

type Hash struct {
	Email string
	Hash  string
}

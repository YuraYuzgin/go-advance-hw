package verify

type Request struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Address  string `json:"address"`
}

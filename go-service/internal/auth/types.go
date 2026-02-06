package auth

type LoginRequest struct {
	Username    string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Message string `json:"message"`
	User    any    `json:"user"`
}

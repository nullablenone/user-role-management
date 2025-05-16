package auth

type RegisterRequest struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
}

type ResponseRegister struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type ResponseLogin struct {
	Token string `json:"token"`
}

package auth

type RegisterParams struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,gte=8"`
	Fullname string `json:"fullname" binding:"required"`
	Username string `json:"username" binding:"required"`
}

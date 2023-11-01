package dtohttp

type CreateAccountRequest struct {
	Username string `json:"username" binding:"required"`
	FullName string `json:"full_name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type CreateAccountResponse struct {
	Username string `json:"username"`
	FullName string `json:"full_name"`
	Email    string `json:"email"`
}

type CheckPasswordRequest struct {
	Password string `json:"password"`
}

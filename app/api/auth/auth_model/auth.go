package auth_model

import "todoBackend/utils/responses"

type LoginRequest struct {
	Username     string `json:"username" binding:"required"` // 用户名，非空且唯一
	PasswordHash string `json:"password"`
}
type LoginResponse struct {
	Token string `json:"jwts"`
}
type RegisterRequest struct {
	Username     string `json:"username"` // 用户名，非空且唯一
	PasswordHash string `json:"password"`
}
type RegisterResponse struct {
	responses.Response
}

type PasswordRequest struct {
	PasswordHash string `json:"password"`
}

package auth

import "github.com/google/uuid"

// Token
// Structure that user sees
// UserID and ID are hidden from user
type Token struct {
	ID           uuid.UUID `json:"id"`
	UserID       int       `json:"user_id"`
	AccessToken  string    `json:"access-token"`
	RefreshToken string    `json:"refresh-token"`
	AccessExp    int64     `json:"refresh-token-exp"`
	RefreshExp   int64     `json:"access-token-exp"`
}

// Login
// used for retrieving login credentials from user
type Login struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type TokenType string

var (
	AccessToken  TokenType = "ACCESS_TOKEN"
	RefreshToken TokenType = "REFRESH_TOKEN"
)

// Payload
// is stored in context
// middleware.Auth() saves structure in the context, after successful authorization
type Payload struct {
	ID   int       `json:"id"`
	Role string    `json:"role"`
	Type TokenType `json:"type"`
	Exp  int64     `json:"exp"`
}

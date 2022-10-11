package auth

type Token struct {
	AccessToken  string `json:"access-token"`
	RefreshToken string `json:"refresh-token"`
	Exp          string `json:"exp"`
}

type Login struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type Payload struct {
	ID   int    `json:"id"`
	Role string `json:"role"`
	Exp  int64  `json:"exp"`
}

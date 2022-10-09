package jwt

type Payload struct {
	ID   int    `json:"id"`
	Role string `json:"role"`
	Exp  int64  `json:"exp"`
}

func (p Payload) Valid() error {
	return nil
}

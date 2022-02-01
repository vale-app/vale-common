package models

type JWTPayload struct {
	UserID    string   `json:"user_id"`
	Roles     []string `json:"roles"`
	FirstName string   `json:"first_name"`
	LastName  string   `json:"last_name"`
	Email     string   `json:"email"`
}

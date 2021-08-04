package model

type Lead struct {
	ID          int64  `json:"id"`
	FirstName   string `json:"firstname"`
	LastName    string `json:"lastname"`
	Email       string `json:"email"`
	Phone       string `json:"phone"`
	Description string `json:"description"`
}
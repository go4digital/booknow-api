package models

import "time"

type Lead struct {
	ID          int       `json:"id" pg:",pk,unique,notnull"`
	FirstName   string    `json:"firstName" pg:",notnull"`
	LastName    string    `json:"lastName" pg:",notnull"`
	Email       string    `json:"email" pg:",notnull"`
	Phone       string    `json:"phone" pg:",notnull"`
	Description string    `json:"description" pg:",notnull"`
	CreatedAt   time.Time `json:"createdAt"`
}

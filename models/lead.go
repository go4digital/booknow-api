package models

import "time"

type Lead struct {
	ID          int       `json:"id" pg:",pk,unique,notnull"`
	FirstName   string    `json:"firstName"`
	LastName    string    `json:"lastName"`
	Email       string    `json:"email" pg:",unique,notnull"`
	Phone       string    `json:"phone"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

type LeadInput struct {
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	Email       string `json:"email" pg:",unique,notnull"`
	Phone       string `json:"phone"`
	Description string `json:"description"`
}

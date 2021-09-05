package models

import "time"

type Lead struct {
	ID          int       `json:"id" bun:",unique,notnull"`
	FirstName   string    `json:"firstName" bun:",notnull"`
	LastName    string    `json:"lastName" bun:",notnull"`
	Email       string    `json:"email" bun:",notnull"`
	Phone       string    `json:"phone" bun:",notnull"`
	Description string    `json:"description" bun:",notnull"`
	CreatedAt   time.Time `json:"createdAt" bun:",notnull,default:current_timestamp"`
}

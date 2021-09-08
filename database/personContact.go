package database

import "github.com/uptrace/bun"

type PersonContact struct {
	bun.BaseModel `bun:"person_contact"`
	Id            int `bun:",unique,notnull"`
	PersonId      int
	Person        []*Person `bun:"rel:has-many"`
	ContactId     int
	Contact       []*Contact `bun:"rel:has-many"`
}

package database

import "github.com/uptrace/bun"

type PersonContact struct {
	bun.BaseModel `bun:"person_contact"`
	Id            int64 `bun:",pk,unique,notnull"`
	PersonId      int64
	Person        *Person `bun:"rel:has-one"`
	ContactId     int64
	Contact       *Contact `bun:"rel:has-one"`
}

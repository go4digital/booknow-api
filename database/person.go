package database

import (
	"time"

	"github.com/uptrace/bun"
)

type Person struct {
	bun.BaseModel `bun:"person"`
	Id            int64     `bun:",unique,notnull"`
	FirstName     string    `bun:",notnull,type:varchar(20)"`
	LastName      string    `bun:",notnull,type:varchar(20)"`
	CreatedAt     time.Time `bun:",nullzero,notnull,default:current_timestamp"`
	CreatedBy     int64
	UpdatedAt     time.Time
	UpdatedBy     int64
	Contacts      []Contact `bun:"m2m:person_contact"`
}

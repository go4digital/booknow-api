package database

import (
	"time"

	"github.com/uptrace/bun"
)

type Company struct {
	bun.BaseModel `bun:"company"`
	Id            int64     `bun:",pk,unique,notnull"`
	Name          string    `bun:",notnull,type:varchar(50)"`
	CreatedAt     time.Time `bun:",nullzero,notnull,default:current_timestamp"`
	CreatedBy     int64
	UpdatedAt     time.Time
	UpdatedBy     int64
	Contacts      []*Contact `bun:"rel:has-many"`
	Persons       []*Person  `bun:"rel:has-many"`
}

package database

import "github.com/uptrace/bun"

type Contact struct {
	bun.BaseModel `bun:"contact"`
	Id            int    `bun:",unique,notnull"`
	Description   string `bun:",notnull,type:varchar(100)"`
	ReferenceId   int
	Reference     *Reference `bun:"rel:has-one"`
}

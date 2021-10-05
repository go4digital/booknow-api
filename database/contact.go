package database

import "github.com/uptrace/bun"

type Contact struct {
	bun.BaseModel `bun:"contact"`
	Id            int64  `bun:",unique,notnull"`
	Description   string `bun:",notnull,type:varchar(100)"`
	ReferenceId   int64
	Reference     *Reference `bun:"rel:has-one"`
}

package database

import "github.com/uptrace/bun"

type Document struct {
	bun.BaseModel `bun:"document"`
	Id            int64 `bun:",unique,notnull"`
	ReferenceId   int64
	Url           string `bun:",notnull,type:varchar(100)"`
}

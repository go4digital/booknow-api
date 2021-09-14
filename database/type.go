package database

import "github.com/uptrace/bun"

type Type struct {
	bun.BaseModel `bun:"type"`
	Id            int64  `bun:",unique,notnull"`
	Description   string `bun:",notnull,type:varchar(45)"`
}

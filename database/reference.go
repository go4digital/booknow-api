package database

import "github.com/uptrace/bun"

type Reference struct {
	bun.BaseModel `bun:"reference"`
	Id            int64  `bun:",unique,notnull"`
	Description   string `bun:",notnull,type:varchar(45)"`
	TypeId        int64
	Type          *Type `bun:"rel:has-one"`
}

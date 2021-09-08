package database

import "github.com/uptrace/bun"

type Reference struct {
	bun.BaseModel `bun:"reference"`
	Id            int    `bun:",unique,notnull"`
	Description   string `bun:",notnull,type:varchar(45)"`
	TypeId        int
	Type          *Type `bun:"rel:has-one"`
}

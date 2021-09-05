package models

type References struct {
	Id          int    `bun:",unique,notnull"`
	Description string `bun:",notnull,type:varchar(45)"`
	TypeId      int
	Type        *Type `bun:"rel:has-one"`
}

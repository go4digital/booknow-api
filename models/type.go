package models

type Type struct {
	Id          int    `bun:",unique,notnull"`
	Description string `bun:",notnull,type:varchar(45)"`
}

package models

type Message struct {
	Id           int    `bun:",unique,notnull"`
	Description  string `bun:",notnull,type:varchar(250)"`
	FromPersonId int
	ToPersonId   int
	ReferencesId int
	References   *References `bun:"rel:has-one"`
}

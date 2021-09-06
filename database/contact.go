package database

type Contact struct {
	Id           int    `bun:",unique,notnull"`
	Description  string `bun:",notnull,type:varchar(100)"`
	ReferencesId int
	References   *References `bun:"rel:has-one"`
}

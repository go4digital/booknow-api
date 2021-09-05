package models

type PersonContact struct {
	Id        int `bun:",unique,notnull"`
	PersonId  int
	Person    *Person `bun:"rel:has-one"`
	ContactId int
	Contact   []*Contact `bun:"rel:has-many"`
}

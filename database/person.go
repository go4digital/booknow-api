package database

import "github.com/uptrace/bun"

type Person struct {
	bun.BaseModel `bun:"person"`
	Id            int    `bun:",unique,notnull"`
	FirstName     string `bun:",notnull,type:varchar(20)"`
	LastName      string `bun:",notnull,type:varchar(20)"`
	ReferenceId   int
	Reference     *Reference `bun:",rel:has-one"`
	TenantId      int
}

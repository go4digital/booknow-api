package models

import "github.com/uptrace/bun"

type Person struct {
	bun.BaseModel `bun:"persons"`
	Id            int    `bun:",unique,notnull"`
	FirstName     string `bun:",notnull,type:varchar(20)"`
	LastName      string `bun:",notnull,type:varchar(20)"`
	ReferenceId   int
	References    *References `bun:",rel:has-one"`
	TenantId      int
}

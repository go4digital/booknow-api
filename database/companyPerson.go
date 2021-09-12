package database

import (
	"github.com/uptrace/bun"
)

type CompanyPerson struct {
	bun.BaseModel `bun:"company_person"`
	Id            int64 `bun:",pk,unique,notnull"`
	CompanyId     int64
	Company       *Company `bun:"rel:has-one"`
	PersonId      int64
	Person        []*Person `bun:"rel:has-many"`
}

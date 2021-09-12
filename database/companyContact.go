package database

import (
	"github.com/uptrace/bun"
)

type CompanyContact struct {
	bun.BaseModel `bun:"company_contact"`
	Id            int64 `bun:",pk,unique,notnull"`
	CompanyId     int64
	Company       *Company `bun:"rel:has-one"`
	ContactId     int64
	Contact       []*Contact `bun:"rel:has-many"`
}

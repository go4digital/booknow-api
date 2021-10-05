package database

import "github.com/uptrace/bun"

type CompanyDocument struct {
	bun.BaseModel `bun:"company_document"`
	Id            int64 `bun:",unique,notnull"`
	CompanyId     int64
	Company       *Company `bun:"rel:belongs-to"`
	DocumentId    int64
	Document      *Document `bun:"rel:belongs-to"`
}

package database

import "github.com/uptrace/bun"

type PersonDocument struct {
	bun.BaseModel `bun:"company_document"`
	Id            int64 `bun:",unique,notnull"`
	PersonId      int64
	Person        *Person `bun:"rel:belongs-to"`
	DocumentId    int64
	Document      *Document `bun:"rel:belongs-to"`
}

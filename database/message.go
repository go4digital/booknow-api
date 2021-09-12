package database

import (
	"time"

	"github.com/uptrace/bun"
)

type Message struct {
	bun.BaseModel `bun:"message"`
	Id            int64  `bun:",pk,unique,notnull"`
	Description   string `bun:",notnull,type:varchar(250)"`
	FromPersonId  int64
	FromPerson    *Person `bun:"rel:has-one,join:from_person_id=id"`
	ToPersonId    int64
	ToPerson      *Person `bun:"rel:has-one,join:to_person_id=id"`
	ReferenceId   int64
	Reference     *Reference `bun:"rel:has-one"`
	CreatedAt     time.Time  `bun:",nullzero,notnull,default:current_timestamp"`
	CreatedBy     int64
	UpdatedAt     time.Time
	UpdatedBy     int64
}

package database

import (
	"time"

	"github.com/uptrace/bun"
)

type Message struct {
	bun.BaseModel `bun:"message"`
	Id            int    `bun:",unique,notnull"`
	Description   string `bun:",notnull,type:varchar(250)"`
	FromPersonId  int
	FromPerson    *Person `bun:"rel:has-one,join:from_person_id=id"`
	ToPersonId    int
	ToPerson      *Person `bun:"rel:has-one,join:to_person_id=id"`
	ReferenceId   int
	Reference     *Reference `bun:"rel:has-one"`
	CreatedAt     time.Time  `bun:",nullzero,notnull,default:current_timestamp"`
	CreatedBy     int
	UpdatedAt     bun.NullTime
	UpdatedBy     int
}

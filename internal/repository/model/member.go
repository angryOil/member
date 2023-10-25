package model

import (
	"github.com/uptrace/bun"
	"time"
)

type Member struct {
	bun.BaseModel `bun:"table:member,alias:m"`

	Id       int       `bun:"id,pk,autoincrement"`
	CafeIid  int       `bun:"cafe_id,notnull"`
	UserIid  int       `bun:"user_id,notnull"`
	Nickname string    `bun:"nickname,notnull"`
	IsBanned bool      `bun:"is_banned,notnull"`
	CreateAt time.Time `bun:"create_at,notnull"`
}

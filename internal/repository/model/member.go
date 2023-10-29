package model

import (
	"github.com/uptrace/bun"
	"member/internal/domain"
	"time"
)

type Member struct {
	bun.BaseModel `bun:"table:members,alias:m"`

	Id        int       `bun:"id,pk,autoincrement"`
	CafeId    int       `bun:"cafe_id,notnull"`
	UserId    int       `bun:"user_id,notnull"`
	Nickname  string    `bun:"nickname,notnull"`
	CreatedAt time.Time `bun:"created_at,notnull"`
}

func ToModel(d domain.MemberDomain) Member {
	return Member{
		Id:        d.Id,
		CafeId:    d.CafeId,
		UserId:    d.UserId,
		Nickname:  d.Nickname,
		CreatedAt: d.CreatedAt,
	}
}

func (m Member) ToDomain() domain.MemberDomain {
	return domain.MemberDomain{
		Id:        m.Id,
		CafeId:    m.CafeId,
		UserId:    m.UserId,
		Nickname:  m.Nickname,
		CreatedAt: m.CreatedAt,
	}
}

func ToDomainList(mList []Member) []domain.MemberDomain {
	domainList := make([]domain.MemberDomain, len(mList))
	for i, m := range mList {
		domainList[i] = m.ToDomain()
	}
	return domainList
}

type CafeIdModel struct {
	bun.BaseModel `bun:"table:members,alias:m"`
	CafeId        int `bun:"cafe_id,notnull"`
}

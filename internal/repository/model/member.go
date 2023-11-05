package model

import (
	"github.com/uptrace/bun"
	"member/internal/domain"
	"member/internal/repository/request"
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

func ToCreateModel(cm request.CreateMember) Member {
	return Member{
		CafeId:    cm.CafeId,
		UserId:    cm.UserId,
		Nickname:  cm.Nickname,
		CreatedAt: cm.CreatedAt,
	}
}

func ToPatchModel(pm request.PatchMember) Member {
	return Member{
		Id:        pm.Id,
		CafeId:    pm.CafeId,
		UserId:    pm.UserId,
		Nickname:  pm.Nickname,
		CreatedAt: pm.CreatedAt,
	}
}

func (m Member) ToDomain() domain.Member {
	return domain.NewMemberBuilder().
		Id(m.Id).
		CafeId(m.CafeId).
		UserId(m.UserId).
		Nickname(m.Nickname).
		CreatedAt(m.CreatedAt).
		Build()
}

func ToDomainList(mList []Member) []domain.Member {
	domainList := make([]domain.Member, len(mList))
	for i, m := range mList {
		domainList[i] = m.ToDomain()
	}
	return domainList
}

type CafeIdModel struct {
	bun.BaseModel `bun:"table:members,alias:m"`
	CafeId        int `bun:"cafe_id,notnull"`
}

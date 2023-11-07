package req

import (
	"member/internal/domain"
	"time"
)

type JoinMemberDto struct {
	Nickname string `json:"nickname"`
}

func (d JoinMemberDto) ToDomain(cafeId, userId int) domain.Member {
	return domain.NewMemberBuilder().
		CafeId(cafeId).
		UserId(userId).
		Nickname(d.Nickname).
		CreatedAt(time.Now()).
		Build()
}

type PatchMemberDto struct {
	Nickname string `json:"nickname"`
}

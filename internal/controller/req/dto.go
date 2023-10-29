package req

import (
	"member/internal/domain"
	"time"
)

type JoinMemberDto struct {
	Nickname string `json:"nickname"`
}

func (d JoinMemberDto) ToDomain(cafeId, userId int) domain.MemberDomain {
	return domain.MemberDomain{
		CafeId:    cafeId,
		UserId:    userId,
		Nickname:  d.Nickname,
		CreatedAt: time.Now(),
	}
}

type PatchMemberDto struct {
	Nickname string `json:"nickname"`
}

func (d PatchMemberDto) ToDomain(cafeId, userId int) domain.MemberDomain {
	return domain.MemberDomain{
		UserId:   userId,
		CafeId:   cafeId,
		Nickname: d.Nickname,
	}
}

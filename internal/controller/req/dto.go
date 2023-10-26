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
		IsBanned:  false,
		CreatedAt: time.Now(),
	}
}

type PatchMemberDto struct {
	MemberId int  `json:"member_id"`
	IsBanned bool `json:"is_banned"`
}

func (d PatchMemberDto) ToDomain(cafeId int) domain.MemberDomain {
	return domain.MemberDomain{
		Id:       d.MemberId,
		CafeId:   cafeId,
		IsBanned: d.IsBanned,
	}
}

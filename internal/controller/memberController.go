package controller

import (
	"context"
	"errors"
	"log"
	"member/internal/controller/req"
	"member/internal/controller/res"
	"member/internal/service"
	page2 "member/page"
)

type MemberController struct {
	s service.MemberService
}

func (c MemberController) RequestJoin(ctx context.Context, dto req.JoinMemberDto, cafeId int, userId int) error {
	d := dto.ToDomain(cafeId, userId)
	err := c.s.RequestJoin(ctx, d)
	return err
}

func (c MemberController) GetJoinCafeIds(ctx context.Context, userId int, reqPage page2.ReqPage) (res.IdTotalCountDto, error) {
	cafeIds, count, err := c.s.GetJoinCafeIds(ctx, userId, reqPage)
	if err != nil {
		log.Println("getCount fail err: ", err)
		return res.IdTotalCountDto{}, errors.New("internal server error")
	}
	return res.NewIdTotalCountDto(cafeIds, count), nil
}

func (c MemberController) GetMemberInfo(ctx context.Context, cafeId int, userId int) (res.MemberInfoDto, error) {
	md, err := c.s.GetMemberInfo(ctx, cafeId, userId)
	if err != nil {
		return res.MemberInfoDto{}, err
	}
	dto := res.ToMemberInfoDto(md)
	return dto, nil

}

func (c MemberController) GetMemberList(ctx context.Context, cafeId int, isBanned bool, reqPage page2.ReqPage) (res.MemberInfoListCountDto, error) {
	mDomainList, count, err := c.s.GetMemberList(ctx, cafeId, isBanned, reqPage)
	if err != nil {
		return res.MemberInfoListCountDto{}, err
	}
	return res.NewMemberInfoListCountDto(res.ToMemberInfoList(mDomainList), count), nil
}

func NewMemberController(s service.MemberService) MemberController {
	return MemberController{s: s}
}

package controller

import (
	"context"
	"errors"
	"log"
	"member/internal/controller/req"
	"member/internal/controller/res"
	"member/internal/service"
	req2 "member/internal/service/req"
	res2 "member/internal/service/res"
	page2 "member/page"
)

type MemberController struct {
	s service.MemberService
}

func NewMemberController(s service.MemberService) MemberController {
	return MemberController{s: s}
}

func (c MemberController) CreateMember(ctx context.Context, dto req.JoinMemberDto, cafeId int, userId int) error {
	err := c.s.CreateMember(ctx, req2.CreateMember{
		CafeId:   cafeId,
		UserId:   userId,
		Nickname: dto.Nickname,
	})
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
	info, err := c.s.GetMemberInfo(ctx, cafeId, userId)
	if err != nil {
		return res.MemberInfoDto{}, err
	}
	return infoToDto(info), nil
}

func infoToDto(info res2.GetMemberInfo) res.MemberInfoDto {
	return res.MemberInfoDto{
		Id:        info.Id,
		UserId:    info.UserId,
		NickName:  info.Nickname,
		CreatedAt: info.CreatedAt,
	}
}

func (c MemberController) GetMemberList(ctx context.Context, cafeId int, reqPage page2.ReqPage) (res.MemberInfoListCountDto, error) {
	mDomainList, count, err := c.s.GetMemberList(ctx, cafeId, reqPage)
	if err != nil {
		return res.MemberInfoListCountDto{}, err
	}
	return res.NewMemberInfoListCountDto(res.ToMemberInfoList(mDomainList), count), nil
}

func (c MemberController) PatchMember(ctx context.Context, id int, dto req.PatchMemberDto) error {
	err := c.s.PatchMember(ctx, req2.PatchMember{
		Id:       id,
		Nickname: dto.Nickname,
	})
	return err
}

func (c MemberController) GetInfoByMemberId(ctx context.Context, memberId int) (res.MemberInfoDto, error) {
	mDomain, err := c.s.GetMemberInfoByMemberCafeId(ctx, memberId)
	if err != nil {
		return res.MemberInfoDto{}, err
	}
	return res.ToMemberInfoDto(mDomain), nil
}

func (c MemberController) GetMemberListByMemberIds(ctx context.Context, idsArr []int) ([]res.MemberInfoDto, error) {
	domains, err := c.s.GetMemberListByMemberIds(ctx, idsArr)
	if err != nil {
		return []res.MemberInfoDto{}, err
	}
	return res.ToMemberInfoList(domains), nil
}

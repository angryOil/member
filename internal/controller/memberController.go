package controller

import (
	"context"
	"member/internal/controller/req"
	"member/internal/service"
)

type MemberController struct {
	s service.MemberService
}

func (c MemberController) RequestJoin(ctx context.Context, dto req.JoinMemberDto, cafeId int, userId int) error {
	d := dto.ToDomain(cafeId, userId)
	err := c.s.RequestJoin(ctx, d)
	return err
}

func NewMemberController(s service.MemberService) MemberController {
	return MemberController{s: s}
}

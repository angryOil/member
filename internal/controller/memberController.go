package controller

import "member/internal/service"

type MemberController struct {
	s service.MemberService
}

func NewMemberController(s service.MemberService) MemberController {
	return MemberController{s: s}
}

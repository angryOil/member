package service

import "member/internal/repository"

type MemberService struct {
	repo repository.MemberRepository
}

func NewMemberService(repo repository.MemberRepository) MemberService {
	return MemberService{repo: repo}
}

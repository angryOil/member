package service

import (
	"context"
	"errors"
	"member/internal/domain"
	"member/internal/repository"
)

type MemberService struct {
	repo repository.MemberRepository
}

func (s MemberService) RequestJoin(ctx context.Context, d domain.MemberDomain) error {
	err := validRequestMember(d)
	if err != nil {
		return err
	}
	err = s.repo.CreateMember(ctx, d)
	return err
}

func validRequestMember(m domain.MemberDomain) error {
	if m.Nickname == "" {
		return errors.New("nickname is empty")
	}
	if m.UserId == 0 {
		return errors.New("invalid user id")
	}
	if m.CafeId == 0 {
		return errors.New("invalid cafe id")
	}
	return nil
}

func NewMemberService(repo repository.MemberRepository) MemberService {
	return MemberService{repo: repo}
}

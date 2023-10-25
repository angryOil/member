package service

import (
	"context"
	"errors"
	"member/internal/domain"
	"member/internal/repository"
	page2 "member/page"
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

func (s MemberService) GetJoinCafeIds(ctx context.Context, userId int, reqPage page2.ReqPage) ([]int, int, error) {
	cafeIds, count, err := s.repo.GetCafeIdsByUserId(ctx, userId, reqPage)
	return cafeIds, count, err
}

func (s MemberService) GetMemberInfo(ctx context.Context, cafeId int, userId int) (domain.MemberDomain, error) {
	if cafeId == 0 {
		return domain.MemberDomain{}, errors.New("invalid cafe id")
	}
	if userId == 0 {
		return domain.MemberDomain{}, errors.New("invalid user id")
	}
	md, err := s.repo.GetMemberInfo(ctx, cafeId, userId)
	return md, err
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

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

func NewMemberService(repo repository.MemberRepository) MemberService {
	return MemberService{repo: repo}
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

// 아래부턴 admin 기능

func (s MemberService) GetMemberList(ctx context.Context, cafeId int, isBanned bool, reqPage page2.ReqPage) ([]domain.MemberDomain, int, error) {
	mDomains, count, err := s.repo.GetMemberList(ctx, cafeId, isBanned, reqPage)
	return mDomains, count, err
}

func (s MemberService) PatchMember(ctx context.Context, d domain.MemberDomain) error {
	err := patchMemberValid(d)
	if err != nil {
		return err
	}

	err = s.repo.PatchMember(ctx, d.Id, d.CafeId,
		func(findDomains []domain.MemberDomain) (domain.MemberDomain, error) {
			if len(findDomains) == 0 {
				return domain.MemberDomain{}, errors.New("no rows error")
			}
			return findDomains[0], nil
		},
		func(m domain.MemberDomain) domain.MemberDomain {
			m.IsBanned = d.IsBanned
			return m
		},
	)
	return err
}

func patchMemberValid(d domain.MemberDomain) error {
	if d.Id == 0 {
		return errors.New("invalid member id")
	}
	if d.CafeId == 0 {
		return errors.New("invalid cafe id")
	}
	return nil
}

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

func (s MemberService) GetMemberList(ctx context.Context, cafeId int, reqPage page2.ReqPage) ([]domain.MemberDomain, int, error) {
	mDomains, count, err := s.repo.GetMemberList(ctx, cafeId, reqPage)
	return mDomains, count, err
}

func (s MemberService) PatchMember(ctx context.Context, d domain.MemberDomain) error {
	err := patchMemberValid(d)
	if err != nil {
		return err
	}

	err = s.repo.PatchMember(ctx, d.CafeId, d.UserId,
		func(findDomains []domain.MemberDomain) (domain.MemberDomain, error) {
			if len(findDomains) == 0 {
				return domain.MemberDomain{}, errors.New("no rows error")
			}
			return findDomains[0], nil
		},
		func(m domain.MemberDomain) domain.MemberDomain {
			m.Nickname = d.Nickname
			return m
		},
	)
	return err
}

func (s MemberService) GetMemberInfoByMemberCafeId(ctx context.Context, memberId int, cafeId int) (domain.MemberDomain, error) {
	mDomain, err := s.repo.GetMemberByMemberCafeId(ctx, memberId, cafeId)
	return mDomain, err
}

func (s MemberService) GetMemberListByMemberIds(ctx context.Context, idsArr []int) ([]domain.MemberDomain, error) {
	domains, err := s.repo.GetMemberListByIds(ctx, idsArr)
	return domains, err
}

func patchMemberValid(d domain.MemberDomain) error {
	if d.CafeId == 0 {
		return errors.New("invalid cafe id")
	}
	if d.UserId == 0 {
		return errors.New("invalid user id")
	}
	if d.Nickname == "" {
		return errors.New("invalid nickname")
	}
	return nil
}

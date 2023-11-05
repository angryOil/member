package service

import (
	"context"
	"errors"
	"member/internal/domain"
	"member/internal/domain/vo"
	"member/internal/repository"
	"member/internal/repository/request"
	"member/internal/service/req"
	"member/internal/service/res"
	page2 "member/page"
	"time"
)

type MemberService struct {
	repo repository.MemberRepository
}

func NewMemberService(repo repository.MemberRepository) MemberService {
	return MemberService{repo: repo}
}

const (
	InvalidUserId = "invalid user id"
	InvalidCafeId = "invalid cafe id"
	NoRowsError   = "no rows error"
)

func (s MemberService) CreateMember(ctx context.Context, cm req.CreateMember) error {
	createdAt := time.Now()
	m := domain.NewMemberBuilder().
		Nickname(cm.Nickname).
		CafeId(cm.CafeId).
		UserId(cm.UserId).
		CreatedAt(createdAt).
		Build()
	err := m.ValidCreate()

	if err != nil {
		return err
	}
	err = s.repo.CreateMember(ctx, request.CreateMember{
		CafeId:    cm.CafeId,
		UserId:    cm.UserId,
		Nickname:  cm.Nickname,
		CreatedAt: createdAt,
	})
	return err
}

func (s MemberService) GetJoinCafeIds(ctx context.Context, userId int, reqPage page2.ReqPage) ([]int, int, error) {
	cafeIds, count, err := s.repo.GetCafeIdsByUserId(ctx, userId, reqPage)
	return cafeIds, count, err
}

func (s MemberService) GetMemberInfo(ctx context.Context, cafeId int, userId int) (res.GetMemberInfo, error) {
	if cafeId < 1 {
		return res.GetMemberInfo{}, errors.New(InvalidCafeId)
	}
	if userId < 1 {
		return res.GetMemberInfo{}, errors.New(InvalidUserId)
	}
	md, err := s.repo.GetMemberInfo(ctx, cafeId, userId)
	if err != nil {
		return res.GetMemberInfo{}, err
	}
	info := voToMemberInfo(md.ToMemberInfo())
	return info, err
}

func voToMemberInfo(vo vo.MemberInfo) res.GetMemberInfo {
	return res.GetMemberInfo{
		Id:        vo.Id,
		UserId:    vo.UserId,
		Nickname:  vo.NickName,
		CreatedAt: convertTimeToString(vo.CreatedAt),
	}
}

var koreaZone, _ = time.LoadLocation("Asia/Seoul")

func convertTimeToString(t time.Time) string {
	if t.IsZero() {
		return ""
	}
	t = t.In(koreaZone)
	return t.Format(time.RFC3339)
}

// 아래부턴 admin 기능

func (s MemberService) GetMemberList(ctx context.Context, cafeId int, reqPage page2.ReqPage) ([]res.GetMemberInfo, int, error) {
	mDomains, count, err := s.repo.GetMemberList(ctx, cafeId, reqPage)
	if err != nil {
		return nil, 0, err
	}
	return domainsToInfoList(mDomains), count, err
}

func domainToInfo(d domain.Member) res.GetMemberInfo {
	return voToMemberInfo(d.ToMemberInfo())
}

func domainsToInfoList(domains []domain.Member) []res.GetMemberInfo {
	results := make([]res.GetMemberInfo, len(domains))
	for i, d := range domains {
		results[i] = voToMemberInfo(d.ToMemberInfo())
	}
	return results
}

func (s MemberService) PatchMember(ctx context.Context, pm req.PatchMember) error {
	var id = pm.Id
	var nickname = pm.Nickname

	err := domain.NewMemberBuilder().
		Id(id).
		Nickname(nickname).
		Build().
		ValidUpdate()

	if err != nil {
		return err
	}

	err = s.repo.PatchMember(ctx, id,
		func(findDomains []domain.Member) (domain.Member, error) {
			if len(findDomains) == 0 {
				return domain.NewMemberBuilder().Build(), errors.New(NoRowsError)
			}
			return findDomains[0], nil
		},
		func(m domain.Member) vo.PatchMember {
			m = m.Patch(nickname)
			return m.ToPatchMember()
		},
	)
	return err
}

func (s MemberService) GetMemberInfoByMemberCafeId(ctx context.Context, memberId int) (res.GetMemberInfo, error) {
	d, err := s.repo.GetMemberByMemberCafeId(ctx, memberId)
	if err != nil {
		return res.GetMemberInfo{}, err
	}
	return domainToInfo(d), nil
}

func (s MemberService) GetMemberListByMemberIds(ctx context.Context, idsArr []int) ([]res.GetMemberInfo, error) {
	domains, err := s.repo.GetMemberListByIds(ctx, idsArr)
	if err != nil {
		return nil, err
	}
	return domainsToInfoList(domains), nil
}

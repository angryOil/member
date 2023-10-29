package repository

import (
	"context"
	"errors"
	"github.com/uptrace/bun"
	"log"
	"member/internal/domain"
	"member/internal/repository/model"
	page2 "member/page"
)

type MemberRepository struct {
	db bun.IDB
}

func NewMemberRepository(db bun.IDB) MemberRepository {
	return MemberRepository{db: db}
}

func (r MemberRepository) CreateMember(ctx context.Context, d domain.MemberDomain) error {
	mModel := model.ToModel(d)
	_, err := r.db.NewInsert().Model(&mModel).Exec(ctx)
	return err
}

func (r MemberRepository) GetCafeIdsByUserId(ctx context.Context, userId int, reqPage page2.ReqPage) ([]int, int, error) {
	var cim []model.CafeIdModel
	query := r.db.NewSelect().Model(&cim).Where("user_id = ?", userId)
	err := query.Limit(reqPage.Size).Offset(reqPage.OffSet).Scan(ctx)
	if err != nil {
		return []int{}, 0, err
	}
	ids := make([]int, len(cim))
	for i, cafeId := range cim {
		ids[i] = cafeId.CafeId
	}
	cnt, err := query.Count(ctx)
	if err != nil {
		return []int{}, 0, err
	}
	return ids, cnt, err
}

func (r MemberRepository) GetMemberInfo(ctx context.Context, cafeId int, userId int) (md domain.MemberDomain, err error) {
	var mModel []model.Member
	defer func() {
		if err != nil {
			log.Println("GetMemberInfo err: ", err)
			err = errors.New("internal server error")
		}
	}()
	err = r.db.NewSelect().Model(&mModel).Where("cafe_id = ? and  user_id = ?", cafeId, userId).Scan(ctx)
	if err != nil {
		return md, err
	}
	if len(mModel) == 0 {
		return md, nil
	}
	md = mModel[0].ToDomain()
	return md, nil
}

func (r MemberRepository) GetMemberList(ctx context.Context, cafeId int, reqPage page2.ReqPage) ([]domain.MemberDomain, int, error) {
	var mModels []model.Member
	query := r.db.NewSelect().Model(&mModels).Where("cafe_id = ?", cafeId)
	err := query.Limit(reqPage.Size).Offset(reqPage.OffSet).Order("id desc").Scan(ctx)
	if err != nil {
		log.Println("GetMemberList scan err: ", err)
		return []domain.MemberDomain{}, 0, errors.New("internal server error")
	}

	cnt, err := query.Count(ctx)
	if err != nil {
		log.Println("GetMemberList count err: ", err)
		return []domain.MemberDomain{}, 0, errors.New("internal server error")
	}

	return model.ToDomainList(mModels), cnt, nil
}

func (r MemberRepository) PatchMember(ctx context.Context, cafeId int, userId int,
	validFindFunc func([]domain.MemberDomain) (domain.MemberDomain, error), // repo 에서 조회한 결과를 validate 함
	mergeFunc func(domain.MemberDomain) domain.MemberDomain) error {
	var findModels []model.Member
	err := r.db.NewSelect().Model(&findModels).Where("cafe_id = ? and user_id = ?", cafeId, userId).Scan(ctx)
	if err != nil {
		log.Println("PatchMember find member err: ", err)
		return errors.New("internal server error")
	}

	domains := model.ToDomainList(findModels)
	validDimain, err := validFindFunc(domains)
	if err != nil {
		return err
	}
	mergedDomain := mergeFunc(validDimain)
	mergedModel := model.ToModel(mergedDomain)
	_, err = r.db.NewInsert().Model(&mergedModel).On("CONFLICT (id) DO UPDATE").Exec(ctx)
	return err
}

func (r MemberRepository) GetMemberByMemberCafeId(ctx context.Context, memberId int, cafeId int) (domain.MemberDomain, error) {
	var mModels []model.Member
	err := r.db.NewSelect().Model(&mModels).Where("id = ? and cafe_id = ?", memberId, cafeId).Scan(ctx)
	if err != nil {
		log.Println("GetMemberByMemberCafeId select err: ", err)
		return domain.MemberDomain{}, errors.New("internal server error")
	}
	if len(mModels) == 0 {
		return domain.MemberDomain{}, nil
	}
	return mModels[0].ToDomain(), nil
}

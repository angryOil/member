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

func (r MemberRepository) GetMemberList(ctx context.Context, cafeId int, isBanned bool, reqPage page2.ReqPage) ([]domain.MemberDomain, int, error) {
	var mModels []model.Member
	query := r.db.NewSelect().Model(&mModels).Where("cafe_id = ? and is_banned = ?", cafeId, isBanned)
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

func NewMemberRepository(db bun.IDB) MemberRepository {
	return MemberRepository{db: db}
}

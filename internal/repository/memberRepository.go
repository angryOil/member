package repository

import (
	"context"
	"github.com/uptrace/bun"
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

func NewMemberRepository(db bun.IDB) MemberRepository {
	return MemberRepository{db: db}
}

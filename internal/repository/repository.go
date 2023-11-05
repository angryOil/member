package repository

import (
	"context"
	"errors"
	"github.com/uptrace/bun"
	"log"
	"member/internal/domain"
	"member/internal/domain/vo"
	"member/internal/repository/model"
	"member/internal/repository/request"
	page2 "member/page"
)

type MemberRepository struct {
	db bun.IDB
}

func NewMemberRepository(db bun.IDB) MemberRepository {
	return MemberRepository{db: db}
}

const (
	InternalServerError = "server internal error"
)

func (r MemberRepository) CreateMember(ctx context.Context, cm request.CreateMember) error {
	mModel := model.ToCreateModel(cm)
	_, err := r.db.NewInsert().Model(&mModel).Exec(ctx)
	return err
}

func (r MemberRepository) GetCafeIdsByUserId(ctx context.Context, userId int, reqPage page2.ReqPage) ([]int, int, error) {
	var cim []model.CafeIdModel
	query := r.db.NewSelect().Model(&cim).Column("cafe_id").Where("user_id = ?", userId)
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

func (r MemberRepository) GetMemberInfo(ctx context.Context, cafeId int, userId int) (md domain.Member, err error) {
	var mModel []model.Member
	defer func() {
		if err != nil {
			log.Println("GetMemberInfo err: ", err)
			err = errors.New(InternalServerError)
		}
	}()
	err = r.db.NewSelect().Model(&mModel).Where("cafe_id = ? and  user_id = ?", cafeId, userId).Scan(ctx)
	if err != nil {
		return nil, err
	}
	if len(mModel) == 0 {
		return domain.NewMemberBuilder().Build(), nil
	}
	md = mModel[0].ToDomain()
	return md, nil
}

func (r MemberRepository) GetMemberList(ctx context.Context, cafeId int, reqPage page2.ReqPage) ([]domain.Member, int, error) {
	var mModels []model.Member
	query := r.db.NewSelect().Model(&mModels).Where("cafe_id = ?", cafeId)
	err := query.Limit(reqPage.Size).Offset(reqPage.OffSet).Order("id desc").Scan(ctx)
	if err != nil {
		log.Println("GetMemberList scan err: ", err)
		return nil, 0, errors.New("internal server error")
	}

	cnt, err := query.Count(ctx)
	if err != nil {
		log.Println("GetMemberList count err: ", err)
		return nil, 0, errors.New("internal server error")
	}

	return model.ToDomainList(mModels), cnt, nil
}

func (r MemberRepository) PatchMember(ctx context.Context, id int,
	validFindFunc func([]domain.Member) (domain.Member, error), // repo 에서 조회한 결과를 validate 함
	mergeFunc func(member domain.Member) vo.PatchMember) error {
	var findModels []model.Member
	db, err := r.db.BeginTx(ctx, nil)

	if err != nil {
		log.Println("PatchMember begin tx err: ", err)
		return errors.New(InternalServerError)
	}
	err = db.NewSelect().Model(&findModels).Where("id = ?", id).Scan(ctx)
	if err != nil {
		log.Println("PatchMember find member err: ", err)
		return errors.New("internal server error")
	}

	domains := model.ToDomainList(findModels)
	validDomain, err := validFindFunc(domains)
	if err != nil {
		return err
	}
	mergedVo := mergeFunc(validDomain)
	m := model.ToPatchModel(request.PatchMember{
		Id:        mergedVo.UserId,
		CafeId:    mergedVo.CafeId,
		UserId:    mergedVo.UserId,
		Nickname:  mergedVo.Nickname,
		CreatedAt: mergedVo.CreatedAt,
	})
	_, err = db.NewInsert().Model(&m).On("CONFLICT (id) DO UPDATE").
		On("conflict (cafe_id,user_id) do update").
		Exec(ctx)
	if err != nil {
		db.Rollback()
		return err
	}
	db.Commit()
	return nil
}

func (r MemberRepository) GetMemberByMemberCafeId(ctx context.Context, memberId int) (domain.Member, error) {
	var mModels []model.Member
	err := r.db.NewSelect().Model(&mModels).Where("id = ?", memberId).Scan(ctx)
	if err != nil {
		log.Println("GetMemberByMemberCafeId select err: ", err)
		return nil, errors.New("internal server error")
	}
	if len(mModels) == 0 {
		return domain.NewMemberBuilder().Build(), nil
	}
	return mModels[0].ToDomain(), nil
}

func (r MemberRepository) GetMemberListByIds(ctx context.Context, idsArr []int) ([]domain.Member, error) {
	var mModels []model.Member
	err := r.db.NewSelect().Model(&mModels).Where("id in (?)", bun.In(idsArr)).Scan(ctx)
	if err != nil {
		log.Println("GetMemberListByIds NewSelect err: ", err)
		return nil, errors.New("internal server error")
	}
	return model.ToDomainList(mModels), nil
}

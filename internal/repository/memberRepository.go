package repository

import (
	"context"
	"github.com/uptrace/bun"
	"member/internal/domain"
	"member/internal/repository/model"
)

type MemberRepository struct {
	db bun.IDB
}

func (r MemberRepository) CreateMember(ctx context.Context, d domain.MemberDomain) error {
	mModel := model.ToModel(d)
	_, err := r.db.NewInsert().Model(&mModel).Exec(ctx)
	return err
}

func NewMemberRepository(db bun.IDB) MemberRepository {
	return MemberRepository{db: db}
}

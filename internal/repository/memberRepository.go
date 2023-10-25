package repository

import "github.com/uptrace/bun"

type MemberRepository struct {
	db bun.IDB
}

func NewMemberRepository(db bun.IDB) MemberRepository {
	return MemberRepository{db: db}
}

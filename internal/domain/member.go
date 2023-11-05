package domain

import (
	"errors"
	"member/internal/domain/vo"
	"time"
)

var _ Member = (*member)(nil)

type Member interface {
	ValidCreate() error
	ValidUpdate() error

	Patch(nickname string) Member
	ToMemberInfo() vo.MemberInfo
	ToPatchMember() vo.PatchMember
}

type member struct {
	id        int
	cafeId    int
	userId    int
	nickname  string
	createdAt time.Time
}

func (m *member) ToPatchMember() vo.PatchMember {
	return vo.PatchMember{
		Id:        m.id,
		CafeId:    m.cafeId,
		UserId:    m.userId,
		Nickname:  m.nickname,
		CreatedAt: m.createdAt,
	}
}

func (m *member) Patch(nickname string) Member {
	m.nickname = nickname
	return m
}

func (m *member) ToMemberInfo() vo.MemberInfo {
	return vo.MemberInfo{
		Id:        m.id,
		UserId:    m.userId,
		NickName:  m.nickname,
		CreatedAt: m.createdAt,
	}
}

const (
	InvalidId       = "invalid member id"
	InvalidUserID   = "invalid user id"
	InvalidCafeId   = "invalid cafe if"
	NicknameIsEmpty = "nickname is empty"
)

func (m *member) ValidCreate() error {
	if m.userId < 1 {
		return errors.New(InvalidUserID)
	}
	if m.cafeId < 1 {
		return errors.New(InvalidCafeId)
	}
	if m.nickname == "" {
		return errors.New(NicknameIsEmpty)
	}
	return nil
}

func (m *member) ValidUpdate() error {
	if m.id < 1 {
		return errors.New(InvalidId)
	}
	if m.nickname == "" {
		return errors.New(NicknameIsEmpty)
	}
	return nil
}

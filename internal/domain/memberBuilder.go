package domain

import "time"

var _ MemberBuilder = (*memberBuilder)(nil)

type MemberBuilder interface {
	Id(id int) MemberBuilder
	CafeId(cafeId int) MemberBuilder
	UserId(userId int) MemberBuilder
	Nickname(nickname string) MemberBuilder
	CreatedAt(createdAt time.Time) MemberBuilder

	Build() Member
}

func NewMemberBuilder() MemberBuilder {
	return &memberBuilder{}
}

type memberBuilder struct {
	id        int
	cafeId    int
	userId    int
	nickname  string
	createdAt time.Time
}

func (m *memberBuilder) Id(id int) MemberBuilder {
	m.id = id
	return m
}

func (m *memberBuilder) CafeId(cafeId int) MemberBuilder {
	m.cafeId = cafeId
	return m
}

func (m *memberBuilder) UserId(userId int) MemberBuilder {
	m.userId = userId
	return m
}

func (m *memberBuilder) Nickname(nickname string) MemberBuilder {
	m.nickname = nickname
	return m
}

func (m *memberBuilder) CreatedAt(createdAt time.Time) MemberBuilder {
	m.createdAt = createdAt
	return m
}

func (m *memberBuilder) Build() Member {
	return &member{
		id:        m.id,
		cafeId:    m.cafeId,
		userId:    m.userId,
		nickname:  m.nickname,
		createdAt: m.createdAt,
	}
}

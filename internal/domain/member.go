package domain

import "time"

type MemberDomain struct {
	Id        int
	CafeId    int
	UserId    int
	Nickname  string
	IsBanned  bool
	CreatedAt time.Time
}

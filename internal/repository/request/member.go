package request

import "time"

type CreateMember struct {
	CafeId    int
	UserId    int
	Nickname  string
	CreatedAt time.Time
}

type PatchMember struct {
	CafeId    int
	Id        int
	UserId    int
	Nickname  string
	CreatedAt time.Time
}

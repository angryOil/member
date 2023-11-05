package vo

import "time"

type PatchMember struct {
	Id        int
	CafeId    int
	UserId    int
	Nickname  string
	CreatedAt time.Time
}

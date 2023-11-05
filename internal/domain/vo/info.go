package vo

import "time"

type MemberInfo struct {
	Id        int
	UserId    int
	NickName  string
	CreatedAt time.Time
}

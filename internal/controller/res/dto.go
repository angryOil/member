package res

import (
	"member/internal/domain"
	"time"
)

type IdTotalCountDto struct {
	Ids   []int `json:"ids"`
	Total int   `json:"total"`
}

func NewIdTotalCountDto(ids []int, total int) IdTotalCountDto {
	return IdTotalCountDto{
		Ids:   ids,
		Total: total,
	}
}

type MemberCafeInfoDto struct {
	MemberId  int    `json:"member_id,omitempty"`
	NickName  string `json:"nick_name,omitempty"`
	CreatedAt string `json:"created_at,omitempty"`
	IsBanned  bool   `json:"is_banned"`
}

func ToMemberCafeInfoDto(d domain.MemberDomain) MemberCafeInfoDto {
	return MemberCafeInfoDto{
		MemberId:  d.Id,
		NickName:  d.Nickname,
		CreatedAt: convertTimeToString(d.CreatedAt),
		IsBanned:  d.IsBanned,
	}
}

var koreaZone, _ = time.LoadLocation("Asia/Seoul")

func convertTimeToString(t time.Time) string {
	if t.IsZero() {
		return ""
	}
	t = t.In(koreaZone)
	return t.Format(time.RFC3339)
}

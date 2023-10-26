package res

import (
	"member/internal/domain"
	"time"
)

// 가입한 cafe id list 페이지용 dto

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

// 해당 카페  member dto

type MemberInfoDto struct {
	MemberId  int    `json:"member_id,omitempty"`
	NickName  string `json:"nick_name,omitempty"`
	CreatedAt string `json:"created_at,omitempty"`
	IsBanned  bool   `json:"is_banned"`
}

func ToMemberInfoList(dList []domain.MemberDomain) []MemberInfoDto {
	infoList := make([]MemberInfoDto, len(dList))
	for i, d := range dList {
		infoList[i] = ToMemberInfoDto(d)
	}
	return infoList
}

type MemberInfoListCountDto struct {
	Members []MemberInfoDto `json:"members"`
	Count   int             `json:"count"`
}

func NewMemberInfoListCountDto(members []MemberInfoDto, total int) MemberInfoListCountDto {
	return MemberInfoListCountDto{
		Members: members,
		Count:   total,
	}
}

func ToMemberInfoDto(d domain.MemberDomain) MemberInfoDto {
	return MemberInfoDto{
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

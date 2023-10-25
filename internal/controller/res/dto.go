package res

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

package utils

type PageRequest struct {
	//页码
	PageNumber int `json:"page_number" form:"page_number"`
	//分页大小
	PageSize int `json:"page_size" form:"page_size"`
}

func NewPageRequest() *PageRequest {
	return &PageRequest{
		PageNumber: 1,
		PageSize:   20,
	}
}

func (r *PageRequest) Offset() int {
	return (r.PageNumber - 1) * r.PageSize
}

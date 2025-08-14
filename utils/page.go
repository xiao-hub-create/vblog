package utils

type PageRequest struct {
	//页码
	PageNumber int `json:"page_number"`
	//分页大小
	PageSize int `json:"page_size"`
}

package response

type PaginationResponse struct {
	Data       any   `json:"data"`
	PageNumber int   `json:"pageNumber"`
	PageSize   int   `json:"pageSize"`
	TotalItem  int64 `json:"totalItem"`
	TotalPage  int64 `json:"totalPage"`
}

func ToPaginationResponse(data any, pageNumber int, pageSize int, totalItem int64) PaginationResponse {
	totalPage := (totalItem + int64(pageSize) - 1) / int64(pageSize)

	return PaginationResponse{
		Data:       data,
		PageNumber: (pageNumber / pageSize) + 1,
		PageSize:   pageSize,
		TotalItem:  totalItem,
		TotalPage:  totalPage,
	}
}

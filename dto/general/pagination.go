package dtogeneral

type PaginationData struct {
	TotalPage   int `json:"total_page,omitempty"`
	TotalItem   int `json:"total_item,omitempty"`
	CurrentPage int `json:"current_page,omitempty"`
	Limit       int `json:"limit,omitempty"`
}

type JSONWithPagination struct {
	Data       any            `json:"data,omitempty"`
	Message    string         `json:"message,omitempty"`
	Pagination PaginationData `json:"pagination,omitempty"`
}

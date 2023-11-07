package dtogeneral

type PaginationData struct {
	TotalPage   int `json:"total_page"`
	TotalItem   int `json:"total_item"`
	CurrentPage int `json:"current_page"`
	Limit       int `json:"limit"`
	// NextPage     string `json:"next_page"`
	// PreviousPage string `json:"previous_page"`
}

type JSONWithPagination struct {
	Data       any            `json:"data,omitempty"`
	Message    string         `json:"message,omitempty"`
	Pagination PaginationData `json:"pagination,omitempty"`
}

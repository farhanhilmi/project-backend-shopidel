package dtogeneral

type JSONResponse struct {
	Data         any    `json:"data,omitempty"`
	Message      string `json:"message,omitempty"`
	AccessToken  string `json:"access_token,omitempty"`
	RefreshToken string `json:"refresh_token,omitempty"`
}

type JSONPagination struct {
	Data       any            `json:"data,omitempty"`
	Message    string         `json:"message,omitempty"`
	Pagination PaginationData `json:"pagination,omitempty"`
}

type ErrResponse struct {
	Error string `json:"error"`
}

type PaginationData struct {
	TotalPage   int `json:"total_page"`
	TotalItem   int `json:"total_item"`
	CurrentPage int `json:"current_page"`
	Limit       int `json:"limit"`
}

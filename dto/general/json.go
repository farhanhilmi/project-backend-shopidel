package dtogeneral

type JSONResponse struct {
	Data         any            `json:"data,omitempty"`
	Message      string         `json:"message,omitempty"`
	AccessToken  string         `json:"access_token,omitempty"`
	RefreshToken string         `json:"refresh_token,omitempty"`
	Pagination   PaginationData `json:"pagination,omitempty"`
}

type ErrResponse struct {
	Error string `json:"error"`
}

type PaginationData struct {
	TotalPage   int `json:"total_page,omitempty"`
	TotalItem   int `json:"total_item,omitempty"`
	CurrentPage int `json:"current_page,omitempty"`
	Limit       int `json:"limit,omitempty"`
}

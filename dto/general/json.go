package dtogeneral

type JSONResponse struct {
	Data         any    `json:"data,omitempty"`
	Message      string `json:"message,omitempty"`
	AccessToken  string `json:"access_token,omitempty"`
	RefreshToken string `json:"refresh_token,omitempty"`
}

type ErrResponse struct {
	Error string `json:"error"`
}

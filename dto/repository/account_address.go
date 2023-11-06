package dtorepository

type AccountAddressRequest struct {
	ID int
}

type AccountAddressResponse struct {
	ID              int
	AccountID       int
	Province        string
	District        string
	SubDistrict     string
	Kelurahan       string
	ZipCode         string
	Detail          string
	IsBuyerDefault  bool
	IsSellerDefault bool
}

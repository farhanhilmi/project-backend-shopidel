package dtorepository

type AccountAddressRequest struct {
	ID        int
	AccountID int
}

type AccountAddressResponse struct {
	ID                   int
	AccountID            int
	Province             string
	District             string
	SubDistrict          string
	Kelurahan            string
	ZipCode              string
	Detail               string
	RajaOngkirDistrictId string
	IsBuyerDefault       bool
	IsSellerDefault      bool
}

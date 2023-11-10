package dtorepository

type AccountAddressRequest struct {
	ID        int
	AccountID int
	SellerID  int
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

type RegisterAddressRequest struct {
	AccountId   int
	ProvinceId  int
	DistrictId  int
	SubDistrict string
	Kelurahan   string
	ZipCode     string
	Detail      string
}

type RegisterAddressResponse struct {
	AccountId   int
	ProvinceId  int
	DistrictId  int
	SubDistrict string
	Kelurahan   string
	ZipCode     string
	Detail      string
}

type UpdateAddressRequest struct {
	AddressId   int
	AccountId   int
	ProvinceId  int
	DistrictId  int
	SubDistrict string
	Kelurahan   string
	ZipCode     string
	Detail      string
}

type UpdateAddressResponse struct {
	AddressId   int
	AccountId   int
	ProvinceId  int
	DistrictId  int
	SubDistrict string
	Kelurahan   string
	ZipCode     string
	Detail      string
}

package dtousecase

type RajaOngkirCost struct {
	Value int    `json:"value"`
	Etd   string `json:"etd"`
	Note  string `json:"note"`
}

type RajaOngkirCourier struct {
	Service     string           `json:"service"`
	Description string           `json:"description"`
	Cost        []RajaOngkirCost `json:"cost"`
}
type DeliveryFeeResponse struct {
	Code  string              `json:"code"`
	Name  string              `json:"name"`
	Costs []RajaOngkirCourier `json:"costs"`
}
type RajaOngkirRes struct {
	Results []DeliveryFeeResponse
}
type RajaOngkirFee struct {
	RajaOngkir RajaOngkirRes `json:"rajaongkir"`
}

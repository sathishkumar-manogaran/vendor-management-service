package models

type Vendors struct {
	Vendors []Vendor `json:Vendors"`
}
type Vendor struct {
	Name     `json:"Name"`
	Services []Service `json:"Services"`
}

type Service struct {
	Name    `json:"name"`
	Country `json:"country"`
}

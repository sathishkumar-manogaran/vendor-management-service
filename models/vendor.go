package models

type Service struct {
	Name    `json:"name"`
	Country `json:"country"`
}

type Vendors struct {
	Name     `json:"Name"`
	Services []Service `json:"Services"`
}

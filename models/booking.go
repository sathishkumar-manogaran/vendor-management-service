package models

type Source struct {
	Address `json:"Address"`
	City    `json:"City"`
	Country `json:"Country"`
}

type ExportCustoms struct {
	Country `json:"Country"`
}

type SourcePort struct {
	City    `json:"City"`
	Country `json:"Country"`
}

type DestinationPort struct {
	City    `json:"City"`
	Country `json:"Country"`
}

type ImportCustoms struct {
	Country `json:"Country"`
}

type Destination struct {
	Address `json:"Address"`
	City    `json:"City"`
	Country `json:"Country"`
}

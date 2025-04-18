package entities

type Price struct {
	Product string  `json:"product"`
	Measure string  `json:"measure"`
	Size    string  `json:"size"`
	Price   float64 `json:"price"`
}

func (p Price) GetKey() string {
	return p.Product + p.Measure + p.Size
}

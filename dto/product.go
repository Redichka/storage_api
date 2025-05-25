package dto

type ProductDTO struct {
	Name      *string  `json:"Name"`
	Price     *float64 `json:"Price"`
	PriceFrom *float64 `json:"CountFrom"`
	PriceTo   *float64 `json:"CountTo"`
}

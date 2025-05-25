package domain

type ProductFilter struct {
	Name      *string
	Price     *float64
	PriceFrom *float64
	PriceTo   *float64
}

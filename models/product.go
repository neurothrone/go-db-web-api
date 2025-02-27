package models

type Product struct {
	Id          int     `json:"id"`
	Title       string  `json:"title"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
	Category    string  `json:"category"`
	Image       string  `json:"image"`
}

// CalculatePrice is a method with a receiver p of type Product.
// It can be called on an instance of Product,
// e.g., product.CalculatePrice().
func (p *Product) CalculatePrice() float64 {
	return p.Price * 1.25
}

// CalculatePrice is a standalone function that takes a Product as an argument.
// It can be called with a Product instance passed as an argument,
// e.g., CalculatePrice(product).
func CalculatePrice(p Product) float64 {
	return p.Price * 1.25
}

package viewmodels

type CartViewModel struct {
	CartID     int                `json:"cart_id"`
	Products   []ProductViewmodel `json:"cart_products"`
	TotalPrice float64            `json:"total_price"`
}

type ProductViewmodel struct {
	ID     string `json:"product_id,omitempty"`
	Name   string `json:"name,omitempty"`
	Price  string `json:"price,omitempty"`
	Amount int    `json:"amount,omitempty"`
}

package models

type (
	// Product
	Product struct {
		ProductId          int    `form:"product_id" json:"product_id"`
		ProductName        string `form:"product_name" json:"product_name"`
		ProductPrice       int    `form:"product_price" json:"product_price"`
		ProductDescription string `form:"product_description" json:"product_description"`
		ProductQuantity    int    `form:"product_quantity" json:"product_quantity"`
	}
)

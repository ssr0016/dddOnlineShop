package products

import "time"

type ProductListResponse struct {
	Id    int    `json:"id"`
	SKU   string `json:"sku"`
	Name  string `json:"name"`
	Stock int16  `json:"stock"`
	Price int    `json:"price"`
}

func NewProductListResponseFromEntity(products []Product) []ProductListResponse {
	var productList []ProductListResponse

	for _, product := range products {
		productList = append(productList, product.ToProductListResponse())
	}

	return productList
}

type ProductDetailResponse struct {
	Id        int       `json:"id"`
	SKU       string    `json:"sku"`
	Name      string    `json:"name"`
	Stock     int16     `json:"stock"`
	Price     int       `json:"price"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

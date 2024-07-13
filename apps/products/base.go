package products

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

func Init(router fiber.Router, db *sqlx.DB) {
	repo := newRepository(db)
	svc := newService(repo)
	handler := newHandler(svc)

	productRoute := router.Group("products")
	{
		productRoute.Post("", handler.CreateHandler)
		productRoute.Get("", handler.GetListProducts)
		productRoute.Get("/sku/:sku", handler.GetProductDetail)
	}
}

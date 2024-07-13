package products

import (
	"net/http"
	infrafiber "onlineShop/infra/fiber"
	"onlineShop/infra/response"

	"github.com/gofiber/fiber/v2"
)

type handler struct {
	svc service
}

func newHandler(svc service) handler {
	return handler{
		svc: svc,
	}
}

func (h handler) CreateHandler(ctx *fiber.Ctx) error {
	var req = CreateProductRequestPayload{}

	if err := ctx.BodyParser(&req); err != nil {
		return infrafiber.NewResponse(
			infrafiber.WithMessage("invalid payload"),
		).Send(ctx)
	}

	if err := h.svc.CreateProduct(ctx.UserContext(), req); err != nil {
		myErr, ok := response.ErrorMapping[err.Error()]
		if !ok {
			myErr = response.ErrorGeneral
		}
		return infrafiber.NewResponse(
			infrafiber.WithMessage("invalid payload"),
			infrafiber.WithError(myErr),
		).Send(ctx)
	}

	return infrafiber.NewResponse(
		infrafiber.WithHttpCode(http.StatusCreated),
		infrafiber.WithMessage("create product success"),
	).Send(ctx)
}

func (h handler) GetListProducts(ctx *fiber.Ctx) error {
	var req = ListProductRequestPayload{}

	if err := ctx.QueryParser(&req); err != nil {
		return infrafiber.NewResponse(
			infrafiber.WithMessage("invalid payload"),
			infrafiber.WithError(response.ErrorBadRequest),
		).Send(ctx)
	}

	products, err := h.svc.ListProducts(ctx.UserContext(), req)
	if err != nil {
		myErr, ok := response.ErrorMapping[err.Error()]
		if !ok {
			myErr = response.ErrorGeneral
		}
		return infrafiber.NewResponse(
			infrafiber.WithMessage(err.Error()),
			infrafiber.WithError(myErr),
		).Send(ctx)
	}

	ProductListResponse := NewProductListResponseFromEntity(products)

	return infrafiber.NewResponse(
		infrafiber.WithHttpCode(http.StatusOK),
		infrafiber.WithMessage("get list products success"),
		infrafiber.WithPayload(ProductListResponse),
		infrafiber.WithQuery(req.GenerateDefaulValue()),
	).Send(ctx)
}

func (h handler) GetProductDetail(ctx *fiber.Ctx) error {
	sku := ctx.Params("sku", "")
	if sku == "" {
		return infrafiber.NewResponse(
			infrafiber.WithMessage("invalid payload"),
			infrafiber.WithError(response.ErrorBadRequest),
		).Send(ctx)
	}

	product, err := h.svc.ProductDetail(ctx.UserContext(), sku)
	if err != nil {
		myErr, ok := response.ErrorMapping[err.Error()]
		if !ok {
			myErr = response.ErrorGeneral
		}
		return infrafiber.NewResponse(
			infrafiber.WithMessage(err.Error()),
			infrafiber.WithError(myErr),
		).Send(ctx)
	}

	productDetail := product.ToProductDetailResponse()

	return infrafiber.NewResponse(
		infrafiber.WithHttpCode(http.StatusOK),
		infrafiber.WithMessage("get product detail  success"),
		infrafiber.WithPayload(productDetail),
	).Send(ctx)
}

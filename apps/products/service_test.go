package products

import (
	"context"
	"log"
	"onlineShop/external/database"
	"onlineShop/infra/response"
	"onlineShop/internal/config"
	"testing"

	"github.com/stretchr/testify/require"
)

var svc service

func init() {
	filename := "../../cmd/api/config.yaml"
	err := config.LoadConfig(filename)
	if err != nil {
		panic(err)
	}

	db, err := database.ConnectPostgres(config.Cfg.DB)
	if err != nil {
		panic(err)
	}

	repo := newRepository(db)
	svc = newService(repo)
}

func TestCreateProduct_Sucess(t *testing.T) {
	req := CreateProductRequestPayload{
		Name:  "test",
		Price: 100,
		Stock: 10,
	}

	err := svc.CreateProduct(context.Background(), req)
	require.Nil(t, err)
}

func TestCreateProduct_Fail(t *testing.T) {
	t.Run("name is required", func(t *testing.T) {
		req := CreateProductRequestPayload{
			Name:  "",
			Price: 1000,
			Stock: 10,
		}
		err := svc.CreateProduct(context.Background(), req)
		require.NotNil(t, err)
		require.Equal(t, response.ErrProductRequired, err)
	})

	t.Run("price is invalid", func(t *testing.T) {
		req := CreateProductRequestPayload{
			Name:  "test",
			Price: 0,
			Stock: 10,
		}
		err := svc.CreateProduct(context.Background(), req)
		require.NotNil(t, err)
		require.Equal(t, response.ErrPriceInvalid, err)
	})

	t.Run("stock is invalid", func(t *testing.T) {
		req := CreateProductRequestPayload{
			Name:  "test",
			Price: 1000,
			Stock: 0,
		}
		err := svc.CreateProduct(context.Background(), req)
		require.NotNil(t, err)
		require.Equal(t, response.ErrStockInvalid, err)
	})
}

func TestListProduct_Success(t *testing.T) {
	pagination := ListProductRequestPayload{
		Cursor: 0,
		Size:   10,
	}
	products, err := svc.ListProducts(context.Background(), pagination)
	require.Nil(t, err)
	require.NotNil(t, products)
	log.Printf("%v\n", products)
}

func TestProductDetail_Success(t *testing.T) {
	req := CreateProductRequestPayload{
		Name:  "test",
		Price: 1000,
		Stock: 10,
	}

	ctx := context.Background()

	err := svc.CreateProduct(ctx, req)
	require.Nil(t, err)

	products, err := svc.ListProducts(ctx, ListProductRequestPayload{
		Cursor: 0,
		Size:   10,
	})

	require.Nil(t, err)
	require.NotNil(t, products)
	require.Greater(t, len(products), 0)

	product, err := svc.ProductDetail(ctx, products[0].SKU)
	require.Nil(t, err)
	require.NotEmpty(t, product)

	log.Printf("%v\n", product)
}

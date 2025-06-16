package product

import "github.com/jetsadawwts/go-microservices/modules/models"

type (
	CreateProductReq struct {
		Title    string  `json:"title" validate:"required,max=64"`
		Price    float64 `json:"price" validate:"required"`
		Damage   int     `json:"damage" validate:"required"`
		ImageUrl string  `json:"image_url" validate:"required,max=255"`
	}
	ProductShowCase struct {
		ProductId string  `json:"product_id"`
		Title     string  `json:"title"`
		Price     float64 `json:"price"`
		Damage    int     `json:"damage"`
		ImageUrl  string  `json:"image_url"`
	}

	ProductSearchReq struct {
		Title string `json:"title"`
		models.PaginateReq
	}

	ProductUpdateReq struct {
		Title    string  `json:"title" validate:"required,max=64"`
		Price    float64 `json:"price" validate:"required"`
		ImageUrl string  `json:"image_url" validate:"required,max=255"`
		Damage   int     `json:"damage" validate:"required"`
	}

	EnableOrDisableProductReq struct {
		UsageStatus bool `json:"usage_status"`
	}
)

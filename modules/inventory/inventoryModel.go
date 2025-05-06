package inventory

import (
	"githib.coom/jetsadawwts/go-microservices/modules/models"
	"githib.coom/jetsadawwts/go-microservices/modules/product"
)

type (
	UpdateInventoryReq struct {
		UserId    string `json:"user_id" validate:"required,max=64"`
		ProductId string `json:"product_id" validate:"required,max=64"`
	}

	ProductInInventory struct {
		InventoryId string `json:"inventory_id"`
		*product.ProductShowCase
	}

	UserInventory struct {
		UserId string `json:"user_id"`
		*models.PaginateRes
	}
)

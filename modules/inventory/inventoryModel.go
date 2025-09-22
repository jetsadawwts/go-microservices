package inventory

import (
	"github.com/jetsadawwts/go-microservices/modules/models"
	"github.com/jetsadawwts/go-microservices/modules/product"
)

type (
	UpdateInventoryReq struct {
		UserId    string `json:"user_id" validate:"required,max=64"`
		ProductId string `json:"product_id" validate:"required,max=64"`
	}

	ProductInInventory struct {
		InventoryId string `json:"inventory_id"`
		UserId string `json:"user_id"`
		*product.ProductShowCase
	}

	InventorySearchReq struct {
		models.PaginateReq
	}


)

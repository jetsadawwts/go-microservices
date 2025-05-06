package payment

type (
	ProductServiceReq struct {
		Product []*ProductServiceReqDatum `json:"products" validate:"required"`
	}

	ProductServiceReqDatum struct {
		ProductId string  `json:"product_id" validate:"required,max=64"`
		Price     float64 `json:"price"`
	}
)

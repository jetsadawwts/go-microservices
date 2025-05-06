package inventory

type (
	Inventory struct {
		Id        string `json:"_id" bson:"_id,omitempty"`
		UserId    string `json:"user_id" bson:"user_id"`
		ProductId string `json:"product_id" bson:"product_id"`
	}
)

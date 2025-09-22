package inventoryUsecase

import (
	"context"
	"fmt"

	"github.com/jetsadawwts/go-microservices/config"
	"github.com/jetsadawwts/go-microservices/modules/inventory"
	"github.com/jetsadawwts/go-microservices/modules/inventory/inventoryRepository"
	"github.com/jetsadawwts/go-microservices/modules/models"
	"github.com/jetsadawwts/go-microservices/modules/product"
	productPb "github.com/jetsadawwts/go-microservices/modules/product/productPb"
	"github.com/jetsadawwts/go-microservices/pkg/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type (
    InventoryUsecaseService interface{
        FindUserProducts(pctx context.Context, cfg *config.Config, userId string, req *inventory.InventorySearchReq) (*models.PaginateRes, error)
    }
    inventoryUsecase        struct {
        inventoryRepository inventoryRepository.InventoryRepositoryService
    }
)

func NewInventoryUsecase(inventoryRepository inventoryRepository.InventoryRepositoryService) InventoryUsecaseService {
    return &inventoryUsecase{
        inventoryRepository: inventoryRepository,
    }
}
func (u *inventoryUsecase) FindUserProducts(pctx context.Context, cfg *config.Config, userId string, req *inventory.InventorySearchReq) (*models.PaginateRes, error) {
	// Filter
    filter := bson.D{}
    if req.Start != "" {
        filter = append(filter, bson.E{"_id", bson.D{{"$gt", utils.ConvertToObjectId(req.Start)}}})
    }
    filter = append(filter, bson.E{"user_id", userId})

    // Option
    opts := make([]*options.FindOptions, 0)
    opts = append(opts, options.Find().SetSort(bson.D{{"_id", 1}}))
    opts = append(opts, options.Find().SetLimit(int64(req.Limit)))

    // Find
    inventoryData, err := u.inventoryRepository.FindUserProducts(pctx, filter, opts)
    if err != nil {
        return nil, err
    }

    productData, err := u.inventoryRepository.FindProductsInIds(pctx, cfg.Grpc.ProductUrl, &productPb.FindProductsInIdsReq{
        Ids: func() []string {
            productIds := make([]string, 0)
            for _, v := range inventoryData {
                productIds = append(productIds, v.ProductId)
            }
            return productIds
        }(),
    })

    if err != nil {
        return nil, err
    }

    productMaps := make(map[string]*product.ProductShowCase)
    for _, v := range productData.Products {
        productMaps[v.Id] = &product.ProductShowCase{
            ProductId:   v.Id,
            Title:    v.Title,
            Price:    v.Price,
            ImageUrl: v.ImageUrl,
            Damage:   int(v.Damage),
        }
    }

    results := make([]*inventory.ProductInInventory, 0)
    for _, v := range inventoryData {
        results = append(results, &inventory.ProductInInventory{
            InventoryId: v.Id.Hex(),
            UserId:      "user:" + v.UserId,
            ProductShowCase: &product.ProductShowCase{
                ProductId: "product:" + v.ProductId,
                Title: productMaps["product:" + v.ProductId].Title,
                Price: productMaps["product:" + v.ProductId].Price,
                Damage: productMaps["product:" + v.ProductId].Damage,
                ImageUrl: productMaps["product:" + v.ProductId].ImageUrl,
            },
        })
    }

    // Count
    total, err := u.inventoryRepository.CountUserProducts(pctx, userId)
    if err != nil {
        return nil, err
    }

    if len(results) == 0 {
        return &models.PaginateRes{
            Data:  make([]*inventory.ProductInInventory, 0),
            Total: total,
            Limit: req.Limit,
            First: models.FirstPaginate{
                Href: fmt.Sprintf("%s/%s?limit=%d", cfg.Paginate.InventoryNextPageBasedUrl, userId, req.Limit),
            },
            Next: models.NextPaginate{
                Start: "",
                Href:  "",
            },
        }, nil
    }

    return &models.PaginateRes{
        Data:  results,
        Total: total,
        Limit: req.Limit,
        First: models.FirstPaginate{
            Href:  fmt.Sprintf("%s/%s?limit=%d", cfg.Paginate.InventoryNextPageBasedUrl, userId, req.Limit),
        },
        Next: models.NextPaginate{
            Start: results[len(results)-1].InventoryId,
            Href:  fmt.Sprintf("%s/%s?limit=%d&start=%s", cfg.Paginate.InventoryNextPageBasedUrl, userId, req.Limit, results[len(results)-1].InventoryId),
        },
    }, nil


}


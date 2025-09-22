package productUsecase

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strings"

	"github.com/jetsadawwts/go-microservices/modules/models"
	"github.com/jetsadawwts/go-microservices/modules/product"
	productPb "github.com/jetsadawwts/go-microservices/modules/product/productPb"
	"github.com/jetsadawwts/go-microservices/modules/product/productRepository"
	"github.com/jetsadawwts/go-microservices/pkg/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type (
	ProductUsecaseService interface{
		CreateProduct(pctx context.Context, req *product.CreateProductReq) (*product.ProductShowCase, error)
		FindOneProduct(pctx context.Context, productId string) (*product.ProductShowCase, error)
		FindManyProducts(pctx context.Context, basePaginateUrl string, req *product.ProductSearchReq) (*models.PaginateRes, error)
		EditProduct(pctx context.Context, productId string, req *product.ProductUpdateReq) (*product.ProductShowCase, error)
		EnableOrDisableProduct(pctx context.Context, productId string) (bool, error)
		FindProductsInIds(pctx context.Context, req *productPb.FindProductsInIdsReq) (*productPb.FindProductsInIdsRes, error) 
	}

	productUsecase struct {
		productRepository productRepository.ProductRepositoryService
	}
)

func NewProductUsecase(productRepository productRepository.ProductRepositoryService) ProductUsecaseService {
	return &productUsecase{
		productRepository: productRepository,
	}
}

func (u *productUsecase) CreateProduct(pctx context.Context, req *product.CreateProductReq) (*product.ProductShowCase, error) {
	if !u.productRepository.IsUniqueProduct(pctx, req.Title) {
		return nil, errors.New("error: this title is already exist")
	}

	productId, err := u.productRepository.InsertOneProduct(pctx, &product.Product{
		Title: req.Title,
		Price: req.Price,
		Damage: req.Damage,
		UsageStatus: true,
		ImageUrl: req.ImageUrl,
		CreatedAt: utils.LocalTime(),
		UpdatedAt: utils.LocalTime(),
	})
	if err != nil {
		return nil, err
	}

	return u.FindOneProduct(pctx, productId.Hex())
}

func (u *productUsecase) FindOneProduct(pctx context.Context, productId string) (*product.ProductShowCase, error) {
	result, err := u.productRepository.FindOneProduct(pctx, productId)

	if err != nil {
		return nil, err
	}

	return &product.ProductShowCase{
		ProductId: "product:" + result.Id.Hex(),
		Title: result.Title,
		Price: result.Price,
		Damage: result.Damage,
		ImageUrl: result.ImageUrl,
	}, nil
}

func (u *productUsecase) FindManyProducts(pctx context.Context, basePaginateUrl string, req *product.ProductSearchReq) (*models.PaginateRes, error) {
	findItemsFilter := bson.D{}
	findItemsOpts := make([]*options.FindOptions, 0)

	countItemsFilter := bson.D{}

	fmt.Printf("req: %+v\n", req)
	
	// Filter
	if req.Start != "" {
		req.Start = strings.TrimPrefix(req.Start, "product:")
		findItemsFilter = append(findItemsFilter, bson.E{"_id", bson.D{{"$gt", utils.ConvertToObjectId(req.Start)}}})
	}

	if req.Title != "" {
		findItemsFilter = append(findItemsFilter, bson.E{"title", primitive.Regex{Pattern: req.Title, Options: "i"}})
		countItemsFilter = append(countItemsFilter, bson.E{"title", primitive.Regex{Pattern: req.Title, Options: "i"}})
	}

	findItemsFilter = append(findItemsFilter, bson.E{"usage_status", true})
	countItemsFilter = append(countItemsFilter, bson.E{"usage_status", true})

	// Options
	findItemsOpts = append(findItemsOpts, options.Find().SetSort(bson.D{{"_id", 1}}))
	findItemsOpts = append(findItemsOpts, options.Find().SetLimit(int64(req.Limit)))

	// Find
	results, err := u.productRepository.FindManyProducts(pctx, findItemsFilter, findItemsOpts)
	if err != nil {
		return nil, err
	}

	// Count
	total, err := u.productRepository.CountProducts(pctx, countItemsFilter)
	if err != nil {
		return nil, err
	}

	if len(results) == 0 {
		return &models.PaginateRes{
			Data:  make([]*product.ProductShowCase, 0),
			Total: total,
			Limit: req.Limit,
			First: models.FirstPaginate{
				Href: fmt.Sprintf("%s?limit=%d&title=%s", basePaginateUrl, req.Limit, req.Title),
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
			Href: fmt.Sprintf("%s?limit=%d&title=%s", basePaginateUrl, req.Limit, req.Title),
		},
		Next: models.NextPaginate{
			Start: results[len(results)-1].ProductId,
			Href:  fmt.Sprintf("%s?limit=%d&title=%s&start=%s", basePaginateUrl, req.Limit, req.Title, results[len(results)-1].ProductId),
		},
	}, nil
}

func (u *productUsecase) EditProduct(pctx context.Context, productId string, req *product.ProductUpdateReq) (*product.ProductShowCase, error) {
	updateData := primitive.M{
		"title":       req.Title,
		"price":       req.Price,
		"damage":      req.Damage,
		"image_url":   req.ImageUrl,
		"updated_at":  utils.LocalTime(),
	}

	if req.Title != "" {
		if !u.productRepository.IsUniqueProduct(pctx, req.Title) {
			log.Printf("Error: EditProduct failed this title is already exist")
			return nil, errors.New("error: this title is already exist")
		}

			updateData["title"] = req.Title
		}

		if req.ImageUrl != "" {
			updateData["image_url"] = req.ImageUrl
		}

		if req.Damage > 0 {
			updateData["damage"] = req.Price
		}
		if req.Price >= 0 {
			updateData["price"] = req.Damage
		}
		updateData["updated_at"] = utils.LocalTime()

	if err := u.productRepository.UpdateOneProduct(pctx, productId, updateData); err != nil {
		return nil, err
	}
	

	return u.FindOneProduct(pctx, productId)
}

func (u *productUsecase) EnableOrDisableProduct(pctx context.Context, productId string) (bool, error) {
	result, err := u.productRepository.FindOneProduct(pctx, productId)
	if err != nil {
		return false,err
	}

	if err := u.productRepository.EnableOrDisableProduct(pctx, productId,  !result.UsageStatus); err != nil {
		return false, err
	}

	return !result.UsageStatus, nil
}

func (u *productUsecase) FindProductsInIds(pctx context.Context, req *productPb.FindProductsInIdsReq) (*productPb.FindProductsInIdsRes, error) {
	filter := bson.D{}

	objectIds := make([]primitive.ObjectID, 0)
	for _, productId := range req.Ids {
		objectIds = append(objectIds, utils.ConvertToObjectId(strings.TrimPrefix(productId, "product:")))
	}

	filter = append(filter, bson.E{"_id", bson.D{{"$in", objectIds}}})
	filter = append(filter, bson.E{"usage_status", true})

	results, err := u.productRepository.FindManyProducts(pctx, filter, nil)
	if err != nil {
		return nil, err
	}

	resultsToRes := make([]*productPb.Product, 0)
	for _, result := range results {
		resultsToRes = append(resultsToRes, &productPb.Product{
			Id: result.ProductId,
			Title: result.Title,
			Price: result.Price,
			Damage: int32(result.Damage),
			ImageUrl: result.ImageUrl,
		})
	}
	
	return &productPb.FindProductsInIdsRes{
		Products: resultsToRes,
	}, nil
}
 
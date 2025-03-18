package api

import (
	"encoding/json"
	"fmt"

	"github.com/24example/toledoapi/models"
	"github.com/24example/toledoapi/toledoapi"
	"github.com/hooklift/gowsdl/soap"
)

type Api struct {
	apiKey string
}

type Method interface {
	GetBrands() ([]models.Brands, error)
	GetProducts(filter *models.ProductFilter) ([]models.Product, error)
	GetProductPrices(filter *models.ProductFilter) ([]models.ProductPrice, error)
}

func NewApi(apiKey string) Method {
	return &Api{
		apiKey: apiKey,
	}
}

func (a *Api) client() toledoapi.ToledoAPIPortType {
	soapClient := soap.NewClient(
		"http://api-ka.toledo24.ru/ka/ws/ToledoAPI",
	)

	toledoClient := toledoapi.NewToledoAPIPortType(soapClient)

	return toledoClient
}

func (a *Api) GetBrands() ([]models.Brands, error) {
	toledoClient := a.client()

	response, err := toledoClient.Brands(&toledoapi.Brands{
		APIKey: a.apiKey,
	})
	if err != nil {
		return nil, err
	}

	var brandsList toledoapi.ResponseObjects[toledoapi.BrandInfo]

	err = json.Unmarshal([]byte(response.Return_), &brandsList)
	if err != nil {
		// Skip invalid date formats but continue processing other products
		if _, ok := err.(*json.UnmarshalTypeError); !ok {
			return nil, fmt.Errorf("failed to unmarshal response: %w", err)
		}
	}

	brands := make([]models.Brands, len(brandsList.Object.Value.Objects))
	for i, brand := range brandsList.Object.Value.Objects {
		brands[i] = models.Brands{
			ID:   brand.Value.Id,
			Name: brand.Value.Name,
		}
	}

	return brands, nil
}

func (a *Api) GetProducts(filter *models.ProductFilter) ([]models.Product, error) {
	toledoClient := a.client()

	strJsonFilter := "{}"
	strPreviousProduct := ""

	if filter.Filter != nil {
		jsonFilter, err := json.Marshal(filter.Filter)
		if err != nil {
			return nil, err
		}

		strJsonFilter = string(jsonFilter)
	}

	if filter.PreviousProduct != nil {
		strPreviousProduct = *filter.PreviousProduct
	}

	productsInfo := toledoapi.ProductsInfo{
		APIKey:          a.apiKey,
		Filter:          &strJsonFilter,
		PreviousProduct: &strPreviousProduct,
	}

	products := make([]models.Product, 0)

	for {
		response, err := toledoClient.ProductsInfo(&productsInfo)
		if err != nil {
			return nil, err
		}

		var productsList toledoapi.ResponseObjects[toledoapi.ProductInfo]

		err = json.Unmarshal([]byte(response.Return_), &productsList)

		if err != nil {
			// Skip invalid date formats but continue processing other products
			if _, ok := err.(*json.UnmarshalTypeError); !ok {
				return nil, fmt.Errorf("failed to unmarshal response: %w", err)
			}
		}

		for _, product := range productsList.Value.Objects {
			productInfo := models.Product{
				ProductId:          product.Value.ProductId,
				Name:               product.Value.Name,
				DeletionMark:       product.Value.DeletionMark,
				Article:            product.Value.Article,
				Brand:              product.Value.Brand,
				BrandId:            product.Value.BrandId,
				ShortName:          product.Value.ShortName,
				CharacteristicId:   product.Value.CharacteristicId,
				CharacteristicName: product.Value.CharacteristicName,
				Unit:               product.Value.Unit,
				GTD:                product.Value.GTD,
				Weight:             product.Value.Weight,
				Volume:             product.Value.Volume,
				Value:              product.Value.Value,
				Group:              product.Value.Group,
				GroupId:            product.Value.GroupId,
				Multiplicity:       product.Value.Multiplicity,
				MinimalPackage:     product.Value.MinimalPackage,
				Country:            product.Value.Country,
				SuppliersMRP:       product.Value.SuppliersMRP,
				Class:              product.Value.Class,
				Description:        product.Value.Description,
				Barcode:            product.Value.Barcode,
				CuttingProduct:     product.Value.CuttingProduct,
				ExternalName:       product.Value.ExternalName,
				Images:             product.Value.Images,
				Features:           nil,
				Analogs:            nil,
			}

			if product.Value.Features != nil {
				productInfo.Features = make([]*models.ProductFeature, len(product.Value.Features))
				for j, feature := range product.Value.Features {
					productInfo.Features[j] = &models.ProductFeature{
						FeatureId:   feature.Value,
						Description: feature.Description,
						Value:       feature.Value,
						ValueType:   feature.ValueType,
					}
				}
			}

			if product.Value.Analogs != nil {
				productInfo.Analogs = make([]*models.ProductAnalog, len(product.Value.Analogs))
				for j, analog := range product.Value.Analogs {
					productInfo.Analogs[j] = &models.ProductAnalog{
						ProductId:          analog.ProductId,
						Name:               analog.Name,
						CharacteristicId:   analog.CharacteristicId,
						CharacteristicName: analog.CharacteristicName,
						Article:            analog.Article,
						Brand:              analog.Brand,
						BrandId:            analog.BrandId,
					}
				}
			}

			products = append(products, productInfo)
		}

		if productsList.Value.LastObject == nil {
			break
		}

		productsInfo.PreviousProduct = &productsList.Value.LastObject.Value
	}

	return products, nil
}

func (a *Api) GetProductPrices(filter *models.ProductFilter) ([]models.ProductPrice, error) {
	toledoClient := a.client()

	strJsonFilter := "{}"
	strPreviousProduct := ""
	onlyPrice := false
	if filter.Filter != nil {
		jsonFilter, err := json.Marshal(filter.Filter)
		if err != nil {
			return nil, err
		}

		strJsonFilter = string(jsonFilter)
	}

	if filter.PreviousProduct != nil {
		strPreviousProduct = *filter.PreviousProduct
	}

	productsInfo := toledoapi.PriceList{
		APIKey:          a.apiKey,
		Filter:          &strJsonFilter,
		PreviousProduct: &strPreviousProduct,
		OnlyPrice:       &onlyPrice,
	}

	products := make([]models.ProductPrice, 0)

	for {
		response, err := toledoClient.PriceList(&productsInfo)
		if err != nil {
			return nil, err
		}

		var productsList toledoapi.ResponseObjects[toledoapi.ProductPrice]

		err = json.Unmarshal([]byte(response.Return_), &productsList)
		if err != nil {
			// Skip invalid date formats but continue processing other products
			if _, ok := err.(*json.UnmarshalTypeError); !ok {
				return nil, fmt.Errorf("failed to unmarshal response: %w", err)
			}
		}

		for _, product := range productsList.Value.Objects {
			productInfo := models.ProductPrice{
				ProductId:          product.Value.ProductId,
				Name:               product.Value.Name,
				DeletionMark:       product.Value.DeletionMark,
				Article:            product.Value.Article,
				Brand:              product.Value.Brand,
				BrandId:            product.Value.BrandId,
				CharacteristicId:   product.Value.CharacteristicId,
				CharacteristicName: product.Value.CharacteristicName,
				Unit:               product.Value.Unit,
				Price:              product.Value.Price,
				Balance:            product.Value.Balance,
				Group:              product.Value.Group,
				GroupId:            product.Value.GroupId,
				SupplierBalance:    product.Value.SupplierBalance,
				OrderExecutionTime: product.Value.OrderExecutionTime.ToGoTime(),
				MinimumBalance:     product.Value.MinimumBalance,
			}

			products = append(products, productInfo)
		}

		if productsList.Value.LastObject == nil {
			break
		}

		productsInfo.PreviousProduct = &productsList.Value.LastObject.Value
	}

	return products, nil

}

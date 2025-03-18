package models

import "time"

type Filter struct {
	Products []string `xml:"Products,omitempty" json:"Products,omitempty"` //  строка. Отбор, по кодам товаров (не более 1000. Не указываем, если не отбираем).
	Groups   []string `xml:"Groups,omitempty" json:"Groups,omitempty"`     //  строка. Отбор, по кодам групп товаров (не более 10. Не указываем, если не отбираем).
	Brands   []string `xml:"Brands,omitempty" json:"Brands,omitempty"`
}

type ProductFilter struct {
	Filter *Filter `xml:"Filter,omitempty" json:"Filter,omitempty"`

	PreviousProduct *string `xml:"PreviousProduct,omitempty" json:"PreviousProduct,omitempty"`
}

type Product struct {
	ProductId          string            `xml:"ProductId,omitempty" json:"ProductId,omitempty"`
	DeletionMark       bool              `xml:"DeletionMark,omitempty" json:"DeletionMark,omitempty"`
	Article            *string           `xml:"Article,omitempty" json:"Article,omitempty"`
	Brand              *string           `xml:"Brand,omitempty" json:"Brand,omitempty"`
	BrandId            *string           `xml:"BrandId,omitempty" json:"BrandId,omitempty"`
	ShortName          string            `xml:"ShortName,omitempty" json:"ShortName,omitempty"`
	Name               *string           `xml:"Name,omitempty" json:"Name,omitempty"`
	CharacteristicId   string            `xml:"CharacteristicId,omitempty" json:"CharacteristicId,omitempty"`
	CharacteristicName *string           `xml:"CharacteristicName,omitempty" json:"CharacteristicName,omitempty"`
	Unit               string            `xml:"Unit,omitempty" json:"Unit,omitempty"`
	GTD                *string           `xml:"GTD,omitempty" json:"GTD,omitempty"`
	Weight             float32           `xml:"Weight,omitempty" json:"Weight,omitempty"`
	Volume             float32           `xml:"Volume,omitempty" json:"Volume,omitempty"`
	Value              float32           `xml:"Value,omitempty" json:"Value,omitempty"`
	Group              string            `xml:"Group,omitempty" json:"Group,omitempty"`
	GroupId            *string           `xml:"GroupId,omitempty" json:"GroupId,omitempty"`
	Multiplicity       float32           `xml:"Multiplicity,omitempty" json:"Multiplicity,omitempty"`
	MinimalPackage     float32           `xml:"MinimalPackage,omitempty" json:"MinimalPackage,omitempty"`
	Country            *string           `xml:"Country,omitempty" json:"Country,omitempty"`
	SuppliersMRP       float32           `xml:"SuppliersMRP,omitempty" json:"SuppliersMRP,omitempty"`
	Class              string            `xml:"Class,omitempty" json:"Class,omitempty"`
	Description        *string           `xml:"Description,omitempty" json:"Description,omitempty"`
	Features           []*ProductFeature `xml:"Features,omitempty" json:"Features,omitempty"`
	Analogs            []*ProductAnalog  `xml:"Analogs,omitempty" json:"Analogs,omitempty"`
	Images             []string          `xml:"Images,omitempty" json:"Images,omitempty"`
	Barcode            *string           `xml:"Barcode,omitempty" json:"Barcode,omitempty"`
	CuttingProduct     *string           `xml:"CuttingProduct,omitempty" json:"CuttingProduct,omitempty"`
	ExternalName       *string           `xml:"ExternalName,omitempty" json:"ExternalName,omitempty"`
}

type ProductFeature struct {
	FeatureId   string `xml:"FeatureId,omitempty" json:"FeatureId,omitempty"`
	Description string `xml:"Description,omitempty" json:"Description,omitempty"`
	Value       string `xml:"Value,omitempty" json:"Value,omitempty"`
	ValueType   string `xml:"ValueType,omitempty" json:"ValueType,omitempty"`
}

type ProductAnalog struct {
	ProductId          string  `xml:"ProductId,omitempty" json:"ProductId,omitempty"`
	Name               string  `xml:"Name,omitempty" json:"Name,omitempty"`
	CharacteristicId   *string `xml:"CharacteristicId,omitempty" json:"CharacteristicId,omitempty"`
	CharacteristicName *string `xml:"CharacteristicName,omitempty" json:"CharacteristicName,omitempty"`
	Article            string  `xml:"Article,omitempty" json:"Article,omitempty"`
	Brand              string  `xml:"Brand,omitempty" json:"Brand,omitempty"`
	BrandId            string  `xml:"BrandId,omitempty" json:"BrandId,omitempty"`
}

type ProductPrice struct {
	ProductId          string    `xml:"ProductId,omitempty" json:"ProductId,omitempty"`
	Name               string    `xml:"Name,omitempty" json:"Name,omitempty"`
	Article            *string   `xml:"Article,omitempty" json:"Article,omitempty"`
	Brand              *string   `xml:"Brand,omitempty" json:"Brand,omitempty"`
	BrandId            *string   `xml:"BrandId,omitempty" json:"BrandId,omitempty"`
	CharacteristicId   *string   `xml:"CharacteristicId,omitempty" json:"CharacteristicId,omitempty"`
	CharacteristicName *string   `xml:"CharacteristicName,omitempty" json:"CharacteristicName,omitempty"`
	Unit               string    `xml:"Unit,omitempty" json:"Unit,omitempty"`
	Price              float32   `xml:"Price,omitempty" json:"Price,omitempty"`
	Balance            float32   `xml:"Balance,omitempty" json:"Balance,omitempty"`
	Group              string    `xml:"Group,omitempty" json:"Group,omitempty"`
	GroupId            *string   `xml:"GroupId,omitempty" json:"GroupId,omitempty"`
	SupplierBalance    float32   `xml:"SupplierBalance,omitempty" json:"SupplierBalance,omitempty"`
	OrderExecutionTime time.Time `xml:"OrderExecutionTime,omitempty" json:"OrderExecutionTime,omitempty"`
	DeletionMark       bool      `xml:"DeletionMark,omitempty" json:"DeletionMark,omitempty"`
	MinimumBalance     float32   `xml:"MinimumBalance,omitempty" json:"MinimumBalance,omitempty"`
}

package main

import (
	"github.com/google/wire"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/srigalamilitan/gogincrud/product"
)

func initProductAPI(db *gorm.DB) product.ProductAPI {
	wire.Build(product.ProvideProductRepository, product.ProviceProductService, product.ProdiveProductAPI)
	return product.ProductAPI{}
}

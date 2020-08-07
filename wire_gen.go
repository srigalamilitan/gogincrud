// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package main

import (
	"github.com/jinzhu/gorm"
	"github.com/srigalamilitan/gogincrud/product"
)

// Injectors from wire.go:

func InitProductAPI(db *gorm.DB) product.ProductAPI {
	productRepository := product.ProvideProductRepository(db)
	productService := product.ProviceProductService(productRepository)
	productAPI := product.ProdiveProductAPI(productService)
	return productAPI
}
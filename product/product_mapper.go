package product

func ToProduct(productDTO ProductDTO) Product {
	return Product{Code: productDTO.Code, Price: productDTO.Price}
}

func ToProductDTO(product Product) ProductDTO {
	return ProductDTO{ID: product.ID, Code: product.Code, Price: product.Price}
}
func ToProductDtos(products []Product) []ProductDTO {
	productDtos := make([]ProductDTO, len(products))
	for i, itm := range products {
		productDtos[i] = ToProductDTO(itm)
	}
	return productDtos

}

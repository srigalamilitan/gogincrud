package product

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProductAPI struct {
	ProductService ProductService
}

func ProdiveProductAPI(p ProductService) ProductAPI {
	return ProductAPI{ProductService: p}
}
func (p *ProductAPI) FindAll(c *gin.Context) {
	products := p.ProductService.FindAll()
	c.JSON(http.StatusOK, gin.H{"products": ToProductDtos(products)})
}
func (p *ProductAPI) FindByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	product := p.ProductService.FindByID(uint(id))

	c.JSON(http.StatusOK, gin.H{"product": ToProductDTO(product)})
}

func (p *ProductAPI) Create(c *gin.Context) {
	var productDTO ProductDTO
	err := c.BindJSON(&productDTO)
	if err != nil {
		log.Println(err)
		c.Status(http.StatusBadRequest)
		return
	}
	createProduct := p.ProductService.Save(ToProduct(productDTO))
	c.JSON(http.StatusOK, gin.H{"product": ToProductDTO(createProduct)})
}

func (p *ProductAPI) Update(c *gin.Context) {
	var productDTO ProductDTO
	err := c.BindJSON(&productDTO)
	if err != nil {
		log.Fatalln(err)
		c.Status(http.StatusBadRequest)
		return
	}
	id, _ := strconv.Atoi(c.Param("id"))
	product := p.ProductService.FindByID(uint(id))
	if product == (Product{}) {
		c.Status(http.StatusBadRequest)
		return
	}
	product.Code = productDTO.Code
	product.Price = productDTO.Price
	p.ProductService.Save(product)
	c.Status(http.StatusOK)
}
func (p *ProductAPI) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	pruduct := p.ProductService.FindByID(uint(id))
	if pruduct == (Product{}) {
		c.Status(http.StatusBadRequest)
		return
	}
	p.ProductService.Delete(pruduct)
	c.Status(http.StatusOK)
}

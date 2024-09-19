package v1

import (
	"flash_sale/pkg/app"
	"flash_sale/pkg/constants"
	"flash_sale/service/product_service"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

// @Summary Add product
// @Tags Product
// @Produce  json
// @Param add_product body app.AddProductData true "Add product request body"
// @Security BearerAuth
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/products [post]
func AddProduct(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		data app.AddProductData
	)

	httpCode, errCode := app.BindAndValid(c, &data)
	if errCode != constants.SUCCESS {
		appG.Response(httpCode, errCode, nil)
		return
	}

	productService := product_service.Product{
		Name:     data.Name,
		Stock:    data.Stock,
		IsActive: data.IsActive,
		Price:    data.Price,
	}
	if err := productService.Add(); err != nil {
		appG.Response(http.StatusInternalServerError, constants.ERROR_ADD_PRODUCT_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, constants.SUCCESS, nil)
}

// @Summary Update product
// @Produce  json
// @Tags Product
// @Param id path int true "Product ID"
// @Param edit_product body app.EditProductData true "Edit product request body"
// @Security BearerAuth
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/products/{id} [put]
func EditProduct(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		data app.EditProductData
	)

	httpCode, errCode := app.BindAndValid(c, &data)
	if errCode != constants.SUCCESS {
		appG.Response(httpCode, errCode, nil)
		return
	}

	productService := product_service.Product{
		ID:       uint(com.StrTo(c.Param("id")).MustInt()),
		Name:     data.Name,
		Stock:    data.Stock,
		IsActive: data.IsActive,
	}

	err := productService.Edit()
	if err != nil {
		appG.Response(http.StatusInternalServerError, constants.ERROR_EDIT_PRODUCT_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, constants.SUCCESS, nil)
}

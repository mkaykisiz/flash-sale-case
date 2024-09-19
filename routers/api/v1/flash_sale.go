package v1

import (
	"encoding/json"
	"flash_sale/pkg/app"
	"flash_sale/pkg/constants"
	"flash_sale/service/flash_sale_service"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"net/http"
)

// @Summary Add Flash Sale
// @Tags FlashSale
// @Produce  json
// @Param add_flash_sale body app.AddFlashSaleRequest true "Add flash sale request body"
// @Security BearerAuth
// @Success 200 {object} app.AddFlashSaleResponse
// @Failure 500 {object} app.Response
// @Router /api/v1/flash-sales [post]
func AddFlashSale(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		data app.AddFlashSaleRequest
	)

	httpCode, errCode := app.BindAndValid(c, &data)
	if errCode != constants.SUCCESS {
		appG.Response(httpCode, errCode, nil)
		return
	}

	flashSaleService := flash_sale_service.FlashSale{
		ProductID:       data.ProductID,
		DiscountPercent: data.DiscountPercent,
		Stock:           data.Stock,
		StartTime:       data.StartTime,
		EndTime:         data.EndTime,
	}
	err, flashSaleID := flashSaleService.Add()
	if err != nil {
		appG.Response(http.StatusInternalServerError, constants.ERROR_ADD_FLASH_SALE_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, constants.SUCCESS, app.AddFlashSaleResponse{ID: flashSaleID})
}

// @Summary Get flash sales
// @Tags FlashSale
// @Produce  json
// @Security BearerAuth
// @Success 200 {object} app.GetFlashSalesResponse
// @Failure 500 {object} app.Response
// @Router /api/v1/flash-sales [get]
func GetFlashSales(c *gin.Context) {
	appG := app.Gin{C: c}

	flashSaleService := flash_sale_service.FlashSale{}

	flashSales, err := flashSaleService.GetAll()
	if err != nil {
		appG.Response(http.StatusInternalServerError, constants.ERROR_GET_FLASH_SALES_FAIL, nil)
		return
	}
	data, err := json.Marshal(flashSales)
	if err != nil {
		appG.Response(http.StatusInternalServerError, constants.ERROR_GET_FLASH_SALES_FAIL, nil)
		return
	}

	var response []app.FlashSale
	err = json.Unmarshal(data, &response)
	if err != nil {
		appG.Response(http.StatusInternalServerError, constants.ERROR_GET_FLASH_SALES_FAIL, nil)
		return
	}
	result := app.GetFlashSalesResponse{
		Count:      len(response),
		FlashSales: response,
	}

	appG.Response(http.StatusOK, constants.SUCCESS, result)
}

// @Summary Buy Flash Sale
// @Tags FlashSale
// @Produce  json
// @Security BearerAuth
// @Param id path int true "ID"
// @Param buy_flash_sale body app.BuyFlashSaleRequest true "Buy flash sale request body"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/flash-sales/{id}/buy [post]
func BuyFlashSale(c *gin.Context) {
	appG := app.Gin{C: c}
	userIDValue, _ := c.Get("user_id")
	data := app.BuyFlashSaleRequest{}
	id := com.StrTo(c.Param("id")).MustInt()
	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message(constants.GetMsg(constants.INVALIDPARAMS))

	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusBadRequest, constants.INVALIDPARAMS, nil)
		return
	}

	httpCode, errCode := app.BindAndValid(c, &data)
	if errCode != constants.SUCCESS {
		appG.Response(httpCode, errCode, nil)
		return
	}
	userID, _ := userIDValue.(uint)
	flashSaleService := flash_sale_service.FlashSale{ID: uint(id), Unit: data.Unit, UserID: userID}
	flashSale, err := flashSaleService.Get()
	if err != nil {
		appG.Response(http.StatusInternalServerError, constants.ERROR_GET_FLASH_SALE_FAIL, nil)
		return
	}
	if flashSale.Stock < data.Unit {
		appG.Response(http.StatusInternalServerError, constants.ERROR_INSUFFICIENT_STOCK, nil)
		return
	}
	err = flashSaleService.Buy()
	if err != nil {
		appG.Response(http.StatusInternalServerError, constants.ERROR_GET_BUY_FAIL, err.Error())
		return
	}

	appG.Response(http.StatusOK, constants.SUCCESS, nil)
}

// @Summary Get a single flash sale
// @Tags FlashSale
// @Produce  json
// @Security BearerAuth
// @Param id path int true "ID"
// @Success 200 {object} app.GetFlashSaleResponse
// @Failure 500 {object} app.Response
// @Router /api/v1/flash-sales/{id} [get]
func GetFlashSale(c *gin.Context) {
	appG := app.Gin{C: c}
	id := com.StrTo(c.Param("id")).MustInt()
	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message(constants.GetMsg(constants.INVALIDPARAMS))

	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusBadRequest, constants.INVALIDPARAMS, nil)
		return
	}

	flashSaleService := flash_sale_service.FlashSale{ID: uint(id)}
	exists, err := flashSaleService.ExistByID()
	if err != nil {
		appG.Response(http.StatusInternalServerError, constants.ERROR_GET_FLASH_SALE_FAIL, nil)
		return
	}
	if !exists {
		appG.Response(http.StatusOK, constants.ERROR_GET_FLASH_SALE_FAIL, nil)
		return
	}

	flashSale, err := flashSaleService.Get()
	if err != nil {
		appG.Response(http.StatusInternalServerError, constants.ERROR_GET_FLASH_SALE_FAIL, nil)
		return
	}

	data, err := json.Marshal(flashSale)
	if err != nil {
		appG.Response(http.StatusInternalServerError, constants.INTERNAL_ERROR, nil)
		return
	}

	var response app.FlashSale
	err = json.Unmarshal(data, &response)
	if err != nil {
		appG.Response(http.StatusInternalServerError, constants.INTERNAL_ERROR, nil)
		return
	}

	appG.Response(http.StatusOK, constants.SUCCESS, response)
}

// @Summary Edit Flash Sale
// @Tags FlashSale
// @Produce  json
// @Security BearerAuth
// @Param id path int true "Flash Sale ID"
// @Param edit_flash_sale body app.EditFlashSaleRequest true "Edit flash sale request body"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/flash-sales/{id} [put]
func EditFlashSale(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		data app.EditFlashSaleRequest
	)

	httpCode, errCode := app.BindAndValid(c, &data)
	if errCode != constants.SUCCESS {
		appG.Response(httpCode, errCode, nil)
		return
	}

	flashSaleService := flash_sale_service.FlashSale{
		ID:              uint(com.StrTo(c.Param("id")).MustInt()),
		Stock:           data.Stock,
		DiscountPercent: data.DiscountPercent,
		StartTime:       data.StartTime,
		EndTime:         data.EndTime,
	}

	err := flashSaleService.Edit()
	if err != nil {
		appG.Response(http.StatusInternalServerError, constants.ERROR_EDIT_FLASH_SALE_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, constants.SUCCESS, nil)
}

// @Summary Delete flash sale
// @Tags FlashSale
// @Produce  json
// @Security BearerAuth
// @Param id path int true "ID"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/flash-sales/{id} [delete]
func DeleteFlashSale(c *gin.Context) {
	appG := app.Gin{C: c}
	id := com.StrTo(c.Param("id")).MustInt()
	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message(constants.GetMsg(constants.INVALIDPARAMS))

	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusOK, constants.INVALIDPARAMS, nil)
		return
	}

	flashSaleService := flash_sale_service.FlashSale{ID: uint(id)}
	exists, err := flashSaleService.ExistByID()
	if err != nil {
		appG.Response(http.StatusInternalServerError, constants.ERROR_GET_FLASH_SALE_FAIL, nil)
		return
	}
	if !exists {
		appG.Response(http.StatusOK, constants.ERROR_GET_FLASH_SALE_FAIL, nil)
		return
	}

	err = flashSaleService.Delete()
	if err != nil {
		appG.Response(http.StatusInternalServerError, constants.ERROR_DELETE_FLASH_SALE_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, constants.SUCCESS, nil)
}

package routers

import (
	"flash_sale/docs"
	"flash_sale/middleware/jwt"
	"flash_sale/routers/api"
	v1 "flash_sale/routers/api/v1"
	"github.com/gin-gonic/gin"

	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

// InitRouter initialize routing information
func InitRouter() *gin.Engine {
	// programmatically set swagger info
	docs.SwaggerInfo.Title = "Flash Sale"
	docs.SwaggerInfo.Description = "Flash Sale Swagger"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8000"
	docs.SwaggerInfo.BasePath = ""
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	r := gin.New()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.POST("/auth", api.GetAuth)

	apiv1 := r.Group("/api/v1")
	apiv1.Use(jwt.JWT())
	{
		apiv1.POST("/products", v1.AddProduct)
		apiv1.PUT("/products/:id", v1.EditProduct)

		apiv1.POST("/flash-sales", v1.AddFlashSale)
		apiv1.GET("/flash-sales", v1.GetFlashSales)
		apiv1.POST("/flash-sales/:id/buy", v1.BuyFlashSale)
		apiv1.GET("/flash-sales/:id", v1.GetFlashSale)
		apiv1.PUT("/flash-sales/:id", v1.EditFlashSale)
		apiv1.DELETE("/flash-sales/:id", v1.DeleteFlashSale)
	}

	return r
}

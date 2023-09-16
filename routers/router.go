package routers

import (
	"basic-trade-api/controllers"
	"basic-trade-api/middlewares"

	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {

	router := gin.Default()

	authRouter := router.Group("/auth")
	{
		authRouter.POST("/register", controllers.UserRegister)
		authRouter.POST("/login", controllers.UserLogin)
	}

	productRouter := router.Group("/products")
	{
		productRouter.GET("/", controllers.GetAllProduct)
		productRouter.GET("/:productUUID", controllers.GetProductByID)

		productRouter.Use(middlewares.Authentication())
		productRouter.POST("/", controllers.CreateProduct)

		productRouter.PUT("/:productUUID", middlewares.ProductAuthorization(), controllers.UpdateProduct)
		productRouter.DELETE("/:productUUID", middlewares.ProductAuthorization(), controllers.DeleteProduct)
	}

	variantRouter := productRouter.Group("/variants")
	{
		variantRouter.GET("/", controllers.GetAllVariant)
		variantRouter.GET("/:variantUUID", controllers.GetVariantByID)

		productRouter.Use(middlewares.Authentication())
		variantRouter.POST("/", middlewares.CreateVariantAuthorization(), controllers.CreateVariant)

		variantRouter.PUT("/:variantUUID", middlewares.VariantAuthorization(), controllers.UpdateVariant)
		variantRouter.DELETE("/:variantUUID", middlewares.VariantAuthorization(), controllers.DeleteVariant)
	}
	return router
}

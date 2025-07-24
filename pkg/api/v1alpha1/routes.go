package v1alpha1

import (
	"github.com/cybercoder/restbill/pkg/api/middleware"
	controllers "github.com/cybercoder/restbill/pkg/api/v1alpha1/controllers/user"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.RouterGroup) {
	cartController := controllers.NewCartController()
	cart := r.Group("/cart")
	cart.Use(middleware.GetUser())
	cart.GET("/", cartController.GetCart)
	cart.POST("/add/:productId", cartController.AddToCart)
}

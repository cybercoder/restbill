package v1alpha1

import "github.com/gin-gonic/gin"

func SetupRoutes(r *gin.RouterGroup) {
	// Resources endpoints
	// resources := r.Group("/resources")
	// {
	// 	resources.GET("", listResources)
	// 	resources.POST("", createResource)
	// }

	// // Additional v1alpha1 specific endpoints
	// r.GET("/experimental", experimentalFeature)
}

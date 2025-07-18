package api

import (
	"github.com/cybercoder/restbill/pkg/api/v1alpha1"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	v1alpha1Group := router.Group("/apis/billing.finance.ik8s.ir/v1alpha1")
	v1alpha1.SetupRoutes(v1alpha1Group)
}

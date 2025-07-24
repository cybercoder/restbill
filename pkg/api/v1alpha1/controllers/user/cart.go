package controllers

import (
	"net/http"

	"github.com/cybercoder/restbill/pkg/services"
	"github.com/cybercoder/restbill/pkg/types"
	"github.com/cybercoder/restbill/pkg/utils"
	"github.com/gin-gonic/gin"
)

type CartController struct {
	cartService *services.CartService
}

func NewCartController() *CartController {
	return &CartController{
		cartService: services.NewCartService(),
	}
}

func (c *CartController) GetCart(ctx *gin.Context) {
	sub := ctx.GetUint("sub")
	cart, err := c.cartService.GetUserCart(sub)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, cart)
}

func (c *CartController) AddToCart(ctx *gin.Context) {
	sub := ctx.GetUint("sub")
	productId, err := utils.StringToUint(ctx.Param("productId"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	body := types.AddProductToCartBody{}
	if err = ctx.ShouldBind(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err, result := c.cartService.AddProductToCart(sub, productId, 1, body.Addons, "IRR")
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, result)
}

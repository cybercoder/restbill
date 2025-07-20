package controllers

import (
	"net/http"

	"github.com/cybercoder/restbill/pkg/services"
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

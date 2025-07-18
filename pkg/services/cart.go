// service/cart_service.go
package service

import (
	"github.com/cybercoder/restbill/pkg/database/models"
	"github.com/cybercoder/restbill/pkg/database/repositories"
)

type CartService struct {
	cartRepo *repositories.Repository[models.Cart]
}

func Constructor() *CartService {
	return &CartService{
		cartRepo: repositories.NewRepository[models.Cart](),
	}
}

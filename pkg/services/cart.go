package services

import (
	"github.com/cybercoder/restbill/pkg/database/models"
	"github.com/cybercoder/restbill/pkg/database/op"
	"github.com/cybercoder/restbill/pkg/database/repositories"
)

type CartService struct {
	cartRepo *repositories.Repository[models.Cart]
}

func NewCartService() *CartService {
	return &CartService{
		cartRepo: repositories.NewRepository[models.Cart](),
	}
}

func (s *CartService) GetUserCart(userId string) (*models.Cart, error) {
	return s.cartRepo.FindFirst([]repositories.Condition{
		{Field: "user_id", Operator: op.Equal, Value: userId},
	})
}

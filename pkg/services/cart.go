package services

import (
	"github.com/cybercoder/restbill/pkg/database/models"
	"github.com/cybercoder/restbill/pkg/database/op"
	"github.com/cybercoder/restbill/pkg/database/repositories"
	"github.com/cybercoder/restbill/pkg/types"
	"github.com/cybercoder/restbill/pkg/utils"
	"github.com/samber/lo"
)

type CartService struct {
	cartRepo    *repositories.Repository[models.Cart]
	addonsRepo  *repositories.Repository[models.Addon]
	productRepo *repositories.Repository[models.Product]
}

func NewCartService() *CartService {
	return &CartService{
		cartRepo:    repositories.NewRepository[models.Cart](),
		productRepo: repositories.NewRepository[models.Product](),
		addonsRepo:  repositories.NewRepository[models.Addon](),
	}
}

func (s *CartService) GetUserCart(userId uint) (*models.Cart, error) {
	return s.cartRepo.FindOrCreate([]repositories.Condition{
		{Field: "user_id", Operator: op.Equal, Value: userId},
	}, models.Cart{UserId: userId}, repositories.QueryOptions{
		Preload: []repositories.Preload{
			{
				Relation: "Items",
			},
		},
	})
}

func (s *CartService) AddProductToCart(userId uint, productId uint, quantity uint, addons []types.Addon, currencyCode string) (error, interface{}) {
	cart, err := s.GetUserCart(userId)
	if err != nil {
		return err, nil
	}

	product, err := s.productRepo.GetByID(productId, repositories.QueryOptions{
		Preload: []repositories.Preload{
			{
				Relation: "Price.Currency",
				Args:     []any{"code=(?)", currencyCode},
			},
			{
				Relation: "Category.Addons.Price.Currency",
				Args:     []any{"code=(?)", currencyCode},
			},
		},
	})
	if err != nil {
		return err, nil
	}

	addonIds := lo.Map(addons, func(addon types.Addon, _ int) uint {
		return addon.ID
	})

	addonRecords, err := s.addonsRepo.FindAll([]repositories.Condition{
		{Field: "id", Operator: op.In, Value: addonIds},
	}, repositories.QueryOptions{
		Preload: []repositories.Preload{
			{
				Relation: "Price.Currency",
				Args:     []any{"code=(?)", currencyCode},
			},
		},
	})
	if err != nil {
		return err, nil
	}

	// find product is in cart or not

	_, index, found := lo.FindIndexOf(cart.Items, func(item models.CartItem) bool {
		return item.ProductId == productId && utils.CompareTwoArraysByIntKey(addons, item.Addons, func(a1 types.Addon) uint { return a1.ID }, func(a2 models.CartItemAddons) uint { return a2.ID })
	})

	if !found {
		cart.Items = append(cart.Items, models.CartItem{
			ProductId: productId,
			Quantity:  quantity,
			Price:     product.Price[0].Amount,
			Addons: lo.Map(addons, func(addon types.Addon, _ int) models.CartItemAddons {
				a, _ := lo.Find(addonRecords, func(a1 *models.Addon) bool {
					return a1.ID == addon.ID
				})

				return models.CartItemAddons{
					AddonId:  a.ID,
					Quantity: addon.Quantity,
					Price:    a.Price[0].Amount,
				}
			}),
		})
	} else {
		cart.Items[index].Quantity += quantity

	}

	// Create new item addons for comparison

	// Save the cart
	_, err = s.cartRepo.Update(cart)
	if err != nil {
		return err, nil
	}

	// Return the response
	return nil, struct {
		Cart    *models.Cart    `json:"cart"`
		Product *models.Product `json:"product"`
		Addons  []*models.Addon `json:"addons"`
	}{
		Cart:    cart,
		Product: product,
		Addons:  addonRecords,
	}
}

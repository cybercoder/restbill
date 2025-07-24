package types

type AddProductToCartBody struct {
	Addons []Addon `json:"addons"`
}

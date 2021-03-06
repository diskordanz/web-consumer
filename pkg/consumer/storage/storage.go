package storage

import (
	"github.com/diskordanz/web-consumer/pkg/consumer/model"
)

type ConsumerStorage interface {
	Get(int) (model.Consumer, error)
	Create(model.Consumer) (model.Consumer, error)
	Update(model.Consumer) (model.Consumer, error)
	Cart(id, count, offset int) ([]model.CartItem, error)
	CreateCartItem(model.CartItem) (model.CartItem, error)
	UpdateCartItem(model.CartItem) (model.CartItem, error)
	GetCartItem(id, productID int) (model.CartItem, error)
	DeleteCartItem(model.CartItem) error
}

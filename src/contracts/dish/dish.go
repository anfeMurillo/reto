package dish

import (
	"context"
	"e-restaurant/models/dish"
)

type Repository interface {
	CreateDish(ctx context.Context, dish *dish.Dish) (*dish.Dish, error)

	GetById(ctx context.Context, dishId int) (*dish.Dish, error)

	GetAll(ctx context.Context) ([]*dish.Dish, error)

	UpdateName(ctx context.Context, dishId int, new string) error

	UpdatePrice(ctx context.Context, dishId int, new string) error
}

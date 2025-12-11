package lostinventory

import (
	"context"
	lostinventory "e-restaurant/models/lostInventory"
	"time"
)

type Repository interface {
	Create(ctx context.Context, lInvenotory *lostinventory.LostInventory) (*lostinventory.LostInventory, error)

	GetByRestaurant(ctx context.Context, restaurantId int, limit int, since time.Time) ([]*lostinventory.LostInventory, error)
}

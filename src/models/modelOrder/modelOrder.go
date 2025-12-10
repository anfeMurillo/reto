package modelorder

import (
	enumorderstatus "e-restaurant/models/enums/enumOrderStatus"
	"time"
)

type Order struct {
	OrderId      int                         `db:"order_id" json:"order_id"`
	UserId       int                         `db:"user_id" json:"user_id"`
	RestaurantId int                         `db:"restaurant_id" json:"restaurant_id"`
	CreatedAt    time.Time                   `db:"created_at" json:"created_at"`
	Status       enumorderstatus.OrderStatus `db:"status" json:"status"`
}

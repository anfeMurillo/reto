package restaurantstaff

import (
	"context"
	restaurantstaff "e-restaurant/models/restaurantStaff"
	userapp "e-restaurant/models/userApp"
)

type Repository interface {
	CreateStaffMember(ctx context.Context, staff *restaurantstaff.RestaurantStaff) (*restaurantstaff.RestaurantStaff, error)

	DeleteStaffMember(ctx context.Context, restaurantId int, userId int) error

	GetStaffByRestaurant(ctx context.Context, restaurantId int, onlyActive bool) ([]*userapp.User, error)
}

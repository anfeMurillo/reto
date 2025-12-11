package payment

import (
	"context"
	paymentstatus "e-restaurant/models/enums/paymentStatus"
	"e-restaurant/models/payment"
)

type Repository interface {
	Create(ctx context.Context, payment *payment.Payment) (*payment.Payment, error)

	GetById(ctx context.Context, paymentId int) (*payment.Payment, error)

	UpdateStatus(ctx context.Context, paymentId int, status paymentstatus.PaymentStatus) error

	UpdateMethod(ctx context.Context, paymentId int, method paymentstatus.PaymentMethod) error
}

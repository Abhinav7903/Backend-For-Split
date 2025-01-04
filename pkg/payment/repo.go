package payment

import "github.com/Abhinav7903/split/factory"

type Repository interface {
	CreatePaymentMethod(paymentMethod *factory.PaymentMethod) error
	GetPaymentMethods(email string) ([]factory.PaymentMethod, error)
	UpdatePaymentMethod(paymentMethod *factory.PaymentMethod) error
	DeletePaymentMethod(paymentType, email string) error
}

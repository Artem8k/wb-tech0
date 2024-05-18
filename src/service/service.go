package service

import (
	"rest-api/src/database"
	"rest-api/src/database/models"
)

type Service interface {
	GetOrder(uid string) models.Order
	GetTestOrder() models.Order
}

type ServiceImpl struct {
	orderRepo database.OrderRepository
}

func New(o database.OrderRepository) *ServiceImpl {
	return &ServiceImpl{
		orderRepo: o,
	}
}

func (s *ServiceImpl) GetOrder(uid string) models.Order {
	return s.orderRepo.Get(uid)
}

func (s *ServiceImpl) GetTestOrder() models.Order {
	return s.orderRepo.GetTest()
}

package usecase

import "rabbitmq-http-publisher/app/application/dto"

type RabbiMqBusUseCase interface {
	Publish(dto dto.Payload) error
}

// type warehouseStateUseCase struct {
// 	repo    repository.WarehouseStateRepository
// 	service *service.WarehouseStateService
// }

// func (u *RabbiMqBusUseCase) GetAvailableCatalogItemQuantity(id string) error {
// 	return u.service.GetAvailableCatalogItemQuantity(id)
// }
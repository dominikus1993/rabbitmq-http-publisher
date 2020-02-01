package usecase

type RabbiMqUseCase interface {
	GetAvailableCatalogItemQuantity(id string) error
}
package usecase

type BaseUseCase[T any] interface {
	SearchBy(by map[string]interface{}) ([]T, error)
	FindAll() ([]T, error)
	FindById(id string) (*T, error)
	SaveData(payload *T) error
	DeleteData(id string) error
}

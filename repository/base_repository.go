package repository

type BaseRepository[T any] interface {
	Search(by map[string]interface{}) ([]T, error)
	List() ([]T, error)
	Get(id string) (*T, error)
	Save(payload *T) error
	Delete(id string) error
}

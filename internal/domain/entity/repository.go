package entity

type Repository[T any] interface {
	FindAll() ([]T, error)
	FindByID(id int64) (*T, error)
	Create(role T) (int64, error)
	Update(role T) (int64, error)
	Delete(id int64) (int64, error)
}

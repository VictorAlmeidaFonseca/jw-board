package adapter

type DatabaseHandler[T any] interface {
	NewPool() (*T, error)
	DestructPool()
}

type DatabaseHandlerAdapter[T any] struct {
	DatabaseHandler[T]
}

func (d *DatabaseHandlerAdapter[T]) NewPool() (*T, error) {
	return d.DatabaseHandler.NewPool()
}

func (d *DatabaseHandlerAdapter[T]) DestructPool() {
	d.DatabaseHandler.DestructPool()
}

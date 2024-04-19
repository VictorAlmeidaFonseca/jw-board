package protocol

type DatabaseAdapter[T any] interface {
	NewPool() (*T, error)
	DestructPool()
}

type DatabaseHandlerAdapter[T any] struct {
	Conn DatabaseAdapter[T]
}

func NewDatabaseHandlerAdapter[T any](conn DatabaseAdapter[T]) DatabaseHandlerAdapter[T] {
	return DatabaseHandlerAdapter[T]{
		Conn: conn,
	}
}

func (d *DatabaseHandlerAdapter[T]) NewPool() (*T, error) {
	return d.Conn.NewPool()
}

func (d *DatabaseHandlerAdapter[T]) DestructPool() {
	d.Conn.DestructPool()
}

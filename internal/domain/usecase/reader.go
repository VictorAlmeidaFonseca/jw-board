package usecase

type Reader interface {
	ReadAll() (interface{}, error)
	ReadByID(id int64) (interface{}, error)
}

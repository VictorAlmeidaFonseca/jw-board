package entity

type Role struct {
	ID   int64
	Name string
}

func (r Role) Get() string {
	return r.Name
}

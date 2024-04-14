package application

import (
	"github.com/VictorAlmeidaFonseca/jw-board/internal/domain/entity"
)

type PersonService struct {
	repository entity.Repository[entity.Person]
}

func NewPersonService(repository entity.Repository[entity.Person]) *PersonService {
	return &PersonService{
		repository: repository,
	}
}

func (s *PersonService) FindAll() ([]entity.Person, error) {
	return s.repository.FindAll()
}

func (s *PersonService) FindByID(id int64) (*entity.Person, error) {
	return s.repository.FindByID(id)
}

func (s *PersonService) Create(person entity.Person) (int64, error) {
	return s.repository.Create(person)
}

func (s *PersonService) Update(person entity.Person) (int64, error) {
	return s.repository.Update(person)
}

func (s *PersonService) Delete(id int64) (int64, error) {
	return s.repository.Delete(id)
}

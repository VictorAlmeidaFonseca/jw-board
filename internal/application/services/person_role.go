package application

import (
	"github.com/VictorAlmeidaFonseca/jw-board/internal/domain/entity"
)

type PersonRoleService struct {
	repository entity.Repository[entity.PersonRole]
}

func NewPersonRoleService(repository entity.Repository[entity.PersonRole]) *PersonRoleService {
	return &PersonRoleService{
		repository: repository,
	}
}

func (s *PersonRoleService) FindAll() ([]entity.PersonRole, error) {
	return s.repository.FindAll()
}

func (s *PersonRoleService) FindByID(id int64) (*entity.PersonRole, error) {
	return s.repository.FindByID(id)
}

func (s *PersonRoleService) Create(personRole entity.PersonRole) (int64, error) {
	return s.repository.Create(personRole)
}

func (s *PersonRoleService) Update(personRole entity.PersonRole) (int64, error) {
	return s.repository.Update(personRole)
}

func (s *PersonRoleService) Delete(id int64) (int64, error) {
	return s.repository.Delete(id)
}

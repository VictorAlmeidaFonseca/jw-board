package application

import (
	"github.com/VictorAlmeidaFonseca/jw-board/internal/domain/entity"

)

type RoleService struct {
	repository entity.Repository[entity.Role]
}

func NewRoleService(repository entity.Repository[entity.Role]) *RoleService {
	return &RoleService{
		repository: repository,
	}
}

func (s *RoleService) FindAll() ([]entity.Role, error) {
	return s.repository.FindAll()
}

func (s *RoleService) FindByID(id int64) (*entity.Role, error) {
	return s.repository.FindByID(id)
}

func (s *RoleService) Create(role entity.Role) (int64, error) {
	return s.repository.Create(role)
}

func (s *RoleService) Update(role entity.Role) (int64, error) {
	return s.repository.Update(role)
}

func (s *RoleService) Delete(id int64) (int64, error) {
	return s.repository.Delete(id)
}

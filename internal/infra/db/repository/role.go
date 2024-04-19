package repository

import (
	"github.com/VictorAlmeidaFonseca/jw-board/internal/domain/entity"
	"github.com/VictorAlmeidaFonseca/jw-board/internal/infra/db/config"
	protocol "github.com/VictorAlmeidaFonseca/jw-board/internal/protocol"
)

type RoleRepository struct {
	poll *config.Pool
}

func NewRoleRepository(adapter *protocol.DatabaseHandlerAdapter[config.Pool]) *RoleRepository {
	return &RoleRepository{
		poll: adapter.Conn.(*config.Pool),
	}
}
func (s *RoleRepository) FindAll() ([]entity.Role, error) {
	var roles []entity.Role
	result := s.poll.Conn.Find(&roles)
	if result.Error != nil {
		return nil, result.Error
	}

	return roles, nil
}

func (s *RoleRepository) FindByID(id int64) (*entity.Role, error) {
	var role entity.Role
	result := s.poll.Conn.First(&role, id)
	if result.Error != nil {
		return nil, result.Error
	}

	return &role, nil
}

func (s *RoleRepository) Create(role entity.Role) (int64, error) {
	result := s.poll.Conn.Create(&role)
	if result.Error != nil {
		return 0, result.Error
	}

	return result.RowsAffected, nil
}

func (s *RoleRepository) Update(role entity.Role) (int64, error) {
	result := s.poll.Conn.Model(&role).Where("id = ?", role.ID).Updates(&role)
	if result.Error != nil {
		return 0, result.Error
	}

	return result.RowsAffected, nil
}

func (s *RoleRepository) Delete(id int64) (int64, error) {
	result := s.poll.Conn.Delete(entity.Role{}, id)
	if result.Error != nil {
		return 0, result.Error
	}

	return result.RowsAffected, nil
}

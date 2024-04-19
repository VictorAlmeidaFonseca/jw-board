package repository

import (
	"github.com/VictorAlmeidaFonseca/jw-board/internal/domain/entity"
	"github.com/VictorAlmeidaFonseca/jw-board/internal/infra/db/config"

	protocol "github.com/VictorAlmeidaFonseca/jw-board/internal/protocol"
)

type PersonRoleRepository struct {
	poll *config.Pool
}

func NewPersonRoleRepository(adapter *protocol.DatabaseHandlerAdapter[config.Pool]) *PersonRoleRepository {
	return &PersonRoleRepository{
		poll: adapter.Conn.(*config.Pool),
	}
}

func (s *PersonRoleRepository) FindAll() ([]entity.PersonRole, error) {
	var PersonRoles []entity.PersonRole
	result := s.poll.Conn.Find(&PersonRoles)
	if result.Error != nil {
		return nil, result.Error
	}

	return PersonRoles, nil
}

func (s *PersonRoleRepository) FindByID(id int64) (*entity.PersonRole, error) {
	var PersonRole entity.PersonRole
	result := s.poll.Conn.First(&PersonRole, id)
	if result.Error != nil {
		return nil, result.Error
	}

	return &PersonRole, nil
}

func (s *PersonRoleRepository) Create(PersonRole entity.PersonRole) (int64, error) {
	result := s.poll.Conn.Create(&PersonRole)
	if result.Error != nil {
		return 0, result.Error
	}

	return result.RowsAffected, nil
}

func (s *PersonRoleRepository) Update(PersonRole entity.PersonRole) (int64, error) {
	result := s.poll.Conn.Model(&PersonRole).Where("id = ?", PersonRole.PersonID).Updates(&PersonRole)
	if result.Error != nil {
		return 0, result.Error
	}

	return result.RowsAffected, nil
}

func (s *PersonRoleRepository) Delete(id int64) (int64, error) {
	result := s.poll.Conn.Delete(entity.PersonRole{}, id)
	if result.Error != nil {
		return 0, result.Error
	}

	return result.RowsAffected, nil
}

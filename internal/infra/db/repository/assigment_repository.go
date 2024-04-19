package repository

import (
	"github.com/VictorAlmeidaFonseca/jw-board/internal/domain/entity"
	"github.com/VictorAlmeidaFonseca/jw-board/internal/infra/db/config"
	protocol "github.com/VictorAlmeidaFonseca/jw-board/internal/protocol"
)

type AssignmentRepository struct {
	poll *config.Pool
}

func NewAssignmentRepository(adapter *protocol.DatabaseHandlerAdapter[config.Pool]) *AssignmentRepository {
	return &AssignmentRepository{
		poll: adapter.Conn.(*config.Pool),
	}
}

func (s *AssignmentRepository) FindAll() ([]entity.Assignment, error) {
	var Assignments []entity.Assignment
	result := s.poll.Conn.Find(&Assignments)
	if result.Error != nil {
		return nil, result.Error
	}

	return Assignments, nil
}

func (s *AssignmentRepository) FindByID(id int64) (*entity.Assignment, error) {
	var Assignment entity.Assignment
	result := s.poll.Conn.First(&Assignment, id)
	if result.Error != nil {
		return nil, result.Error
	}

	return &Assignment, nil
}

func (s *AssignmentRepository) Create(Assignment entity.Assignment) (int64, error) {
	result := s.poll.Conn.Create(&Assignment)
	if result.Error != nil {
		return 0, result.Error
	}

	return result.RowsAffected, nil
}

func (s *AssignmentRepository) Update(Assignment entity.Assignment) (int64, error) {
	result := s.poll.Conn.Model(&Assignment).Where("id = ?", Assignment.ID).Updates(&Assignment)
	if result.Error != nil {
		return 0, result.Error
	}

	return result.RowsAffected, nil
}

func (s *AssignmentRepository) Delete(id int64) (int64, error) {
	result := s.poll.Conn.Delete(entity.Assignment{}, id)
	if result.Error != nil {
		return 0, result.Error
	}

	return result.RowsAffected, nil
}

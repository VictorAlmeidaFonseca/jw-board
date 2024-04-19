package application

import (
	"github.com/VictorAlmeidaFonseca/jw-board/internal/domain/entity"
)

type AssignmentService struct {
	repository entity.Repository[entity.Assignment]
}

func NewAssignmentService(repository entity.Repository[entity.Assignment]) *AssignmentService {
	return &AssignmentService{
		repository: repository,
	}
}

func (s *AssignmentService) FindAll() ([]entity.Assignment, error) {
	return s.repository.FindAll()
}

func (s *AssignmentService) FindByID(id int64) (*entity.Assignment, error) {
	return s.repository.FindByID(id)
}

func (s *AssignmentService) Create(assignment entity.Assignment) (int64, error) {
	return s.repository.Create(assignment)
}

func (s *AssignmentService) Update(assignment entity.Assignment) (int64, error) {
	return s.repository.Update(assignment)
}

func (s *AssignmentService) Delete(id int64) (int64, error) {
	return s.repository.Delete(id)
}

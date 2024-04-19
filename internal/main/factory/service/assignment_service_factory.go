package factory

import (
	application "github.com/VictorAlmeidaFonseca/jw-board/internal/application/services"
	factory "github.com/VictorAlmeidaFonseca/jw-board/internal/main/factory/repository"
)

func NewAssignmentService() *application.AssignmentService {
	repo := factory.AssignmentRepositoryFactory()
	service := application.NewAssignmentService(repo)
	return service
}

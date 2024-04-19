package factory

import (
	repository "github.com/VictorAlmeidaFonseca/jw-board/internal/infra/db/repository"
	factory "github.com/VictorAlmeidaFonseca/jw-board/internal/main/factory/connection"
)

func AssignmentRepositoryFactory() *repository.AssignmentRepository {
	conn := factory.ConnectionFactory()
	repo := repository.NewAssignmentRepository(&conn)
	return repo
}

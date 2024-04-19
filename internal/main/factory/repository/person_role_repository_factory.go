package factory

import (
	repository "github.com/VictorAlmeidaFonseca/jw-board/internal/infra/db/repository"
	factory "github.com/VictorAlmeidaFonseca/jw-board/internal/main/factory/connection"
)

func PersonRoleRepositoryFactory() *repository.PersonRoleRepository {
	conn := factory.ConnectionFactory()
	repo := repository.NewPersonRoleRepository(&conn)
	return repo
}

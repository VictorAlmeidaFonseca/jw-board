package factory

import (
	repository "github.com/VictorAlmeidaFonseca/jw-board/internal/infra/db/repository"
	factory "github.com/VictorAlmeidaFonseca/jw-board/internal/main/factory/connection"
)

func RoleRepositoryFactory() *repository.RoleRepository {
	conn := factory.ConnectionFactory()
	repo := repository.NewRoleRepository(&conn)
	return repo
}

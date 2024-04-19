package factory

import (
	application "github.com/VictorAlmeidaFonseca/jw-board/internal/application/services"
	factory "github.com/VictorAlmeidaFonseca/jw-board/internal/main/factory/repository"
)

func NewPersonRoleService() *application.PersonRoleService {
	repo := factory.PersonRoleRepositoryFactory()
	service := application.NewPersonRoleService(repo)
	return service
}

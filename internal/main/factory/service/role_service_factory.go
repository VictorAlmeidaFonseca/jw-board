package factory

import (
	application "github.com/VictorAlmeidaFonseca/jw-board/internal/application/services"
	factory "github.com/VictorAlmeidaFonseca/jw-board/internal/main/factory/repository"
)

func NewRoleService() *application.RoleService {
	repo := factory.RoleRepositoryFactory()
	service := application.NewRoleService(repo)
	return service
}

package factory

import (
	application "github.com/VictorAlmeidaFonseca/jw-board/internal/application/services"
	factory "github.com/VictorAlmeidaFonseca/jw-board/internal/main/factory/repository"
)

func NewPersonService() *application.PersonService {
	repo := factory.PersonRepositoryFactory()
	service := application.NewPersonService(repo)
	return service
}

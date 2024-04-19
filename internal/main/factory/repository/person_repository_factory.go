package factory

import (
	repository "github.com/VictorAlmeidaFonseca/jw-board/internal/infra/db/repository"
	factory "github.com/VictorAlmeidaFonseca/jw-board/internal/main/factory/connection"
)

func PersonRepositoryFactory() *repository.PersonRepository {
	conn := factory.ConnectionFactory()
	repo := repository.NewPersonRepository(&conn)
	return repo
}

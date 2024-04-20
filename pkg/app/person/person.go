package pkg

import (
	"context"

	application "github.com/VictorAlmeidaFonseca/jw-board/internal/application/services"
	"github.com/VictorAlmeidaFonseca/jw-board/internal/domain/entity"
	factory "github.com/VictorAlmeidaFonseca/jw-board/internal/main/factory/service"
)

type App struct {
	ctx           context.Context
	PersonService *application.PersonService
}

func NewApp() *App {
	return &App{}
}

func (a *App) Startup(ctx context.Context) {
	a.PersonService = factory.NewPersonService()
	a.ctx = ctx
}

func (a *App) FindAll(ctx context.Context) ([]entity.Person, error) {
	return a.PersonService.FindAll()
}

func (a *App) FindByID(ctx context.Context, id int64) (*entity.Person, error) {
	return a.PersonService.FindByID(id)
}

func (a *App) Create(ctx context.Context, person entity.Person) (int64, error) {
	return a.PersonService.Create(person)
}

func (a *App) Update(ctx context.Context, person entity.Person) (int64, error) {
	return a.PersonService.Update(person)
}

func (a *App) Delete(ctx context.Context, id int64) (int64, error) {
	return a.PersonService.Delete(id)
}

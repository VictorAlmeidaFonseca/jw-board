package main

import (
	"context"

	application "github.com/VictorAlmeidaFonseca/jw-board/internal/application/services"
	factory "github.com/VictorAlmeidaFonseca/jw-board/internal/main/factory/service"
)

type App struct {
	ctx               context.Context
	RoleService       *application.RoleService
	PersonRoleService *application.PersonRoleService
	PersonService     *application.PersonService
	AssignmentService *application.AssignmentService
}

func NewApp() *App {
	return &App{}
}

func (a *App) Startup(ctx context.Context) {
	a.RoleService = factory.NewRoleService()
	a.PersonService = factory.NewPersonService()
	a.PersonRoleService = factory.NewPersonRoleService()
	a.AssignmentService = factory.NewAssignmentService()
	a.ctx = ctx
}

func (a *App) Shutdown(ctx context.Context) {}

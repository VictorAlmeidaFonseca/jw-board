package main

import (
	"fmt"
	"log"

	application "github.com/VictorAlmeidaFonseca/jw-board/internal/application/services"
	"github.com/VictorAlmeidaFonseca/jw-board/internal/domain/entity"
	"github.com/VictorAlmeidaFonseca/jw-board/internal/infra/db/repository"
)

func main() {
	repo := repository.NewRoleRepository()
	usecase := application.NewRoleService(repo)

	usecase.Create(entity.Role{Name: "Indicador"})
	// usecase.Update(entity.Role{Name: "Volante"})
	// usecase.Delete(1)

	data, err := usecase.FindAll()
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(data)

}

package main

import (
	"fmt"
	"log"

	application "github.com/VictorAlmeidaFonseca/jw-board/internal/application/services"
	"github.com/VictorAlmeidaFonseca/jw-board/internal/infra/db/repository"
)

func main() {
	repo := repository.NewRoleRepository()
	usecase := application.NewRoleService(repo)
	data, err := usecase.FindAll()
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(data)

}

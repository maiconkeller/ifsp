package main

import (
	"fmt"
	"os"
	"receitas/internal/controller"
	"receitas/internal/controller/api"
	"receitas/internal/entity"
	"receitas/internal/repository"
	"receitas/internal/repository/inmemory"
	"receitas/internal/repository/postgres"
	"receitas/internal/usecase/buscarreceita"
	"receitas/internal/usecase/novareceita"
	"receitas/internal/usecase/removereceita"
	"receitas/internal/usecase/todasreceitas"
)

func main() {
	fmt.Println("Iniciando o sistema de receitas")

	var repository repository.Interactor

	args := os.Args
	if len(args) > 1 && args[1] == "INMEMORY" {
		fmt.Println("executando db inmemory")
		repository = inmemory.New(map[int64]entity.Receita{})
	} else {
		fmt.Println("executando db postgres")
		repository = postgres.New(postgres.NewDB())
	}

	novareceitaUc := novareceita.New(repository)
	removereceitaUc := removereceita.New(repository)
	todasreceitasUc := todasreceitas.New(repository)
	buscarreceitaUC := buscarreceita.New(repository)

	handle := api.New(novareceitaUc, removereceitaUc, todasreceitasUc, buscarreceitaUC)
	controller.Start(":8080", handle)
}

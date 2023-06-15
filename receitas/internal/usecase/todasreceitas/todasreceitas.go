package todasreceitas

import (
	"receitas/internal/entity"
	"receitas/internal/repository"
)

type UseCase interface {
	BuscarReceitas() []entity.Receita
}

func New(repository repository.Interactor) UseCase {
	return &todasReceitas{
		repository: repository,
	}
}

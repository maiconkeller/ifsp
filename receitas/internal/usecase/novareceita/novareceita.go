package novareceita

import (
	"receitas/internal/entity"
	"receitas/internal/repository"
)

type UseCase interface {
	Criar(receita entity.Receita) (entity.Receita, error)
}

func New(repository repository.Interactor) UseCase {
	return &novaReceita{
		repository: repository,
	}
}

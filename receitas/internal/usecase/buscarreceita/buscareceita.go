package buscarreceita

import (
	"receitas/internal/entity"
	"receitas/internal/repository"
)

type UseCase interface {
	Buscar(id int64) (entity.Receita, error)
}

func New(repository repository.Interactor) UseCase {
	return &buscarreceita{
		repository: repository,
	}
}

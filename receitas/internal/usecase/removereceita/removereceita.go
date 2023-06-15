package removereceita

import (
	"receitas/internal/repository"
)

type UseCase interface {
	Remover(id int64) error
}

func New(repository repository.Interactor) UseCase {
	return &removereceita{
		repository: repository,
	}
}

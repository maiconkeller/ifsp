package novareceita

import (
	"errors"
	"receitas/internal/entity"
	"receitas/internal/repository"
)

type novaReceita struct {
	repository repository.Interactor
}

func (n *novaReceita) Criar(receita entity.Receita) (entity.Receita, error) {
	if len(receita.Ingredientes) == 0 {
		return entity.Receita{}, errors.New("os ingredientes são obrigatorios")
	}
	return n.repository.Create(receita)
}

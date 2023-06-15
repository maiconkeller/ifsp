package buscarreceita

import (
	"errors"
	"receitas/internal/entity"
	"receitas/internal/repository"
)

type buscarreceita struct {
	repository repository.Interactor
}

func (r *buscarreceita) Buscar(id int64) (entity.Receita, error) {
	if id <= 0 {
		return entity.Receita{}, errors.New("o registro deve ter ID maior que zero")
	}
	receita, err := r.repository.Find(id)
	if err != nil {
		return entity.Receita{}, err
	}

	return receita, nil
}

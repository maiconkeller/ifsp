package removereceita

import (
	"errors"
	"receitas/internal/repository"
)

type removereceita struct {
	repository repository.Interactor
}

func (r *removereceita) Remover(id int64) error {
	if id <= 0 {
		return errors.New("o registro deve ter ID maior que zero")
	}
	return r.repository.Delete(id)
}

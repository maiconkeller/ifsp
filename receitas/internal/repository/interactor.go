package repository

import (
	"receitas/internal/entity"
)

type Interactor interface {
	Create(receita entity.Receita) (entity.Receita, error)
	Delete(id int64) error
	FindAll() (receitas []entity.Receita, err error)
	Find(id int64) (receita entity.Receita, err error)
}

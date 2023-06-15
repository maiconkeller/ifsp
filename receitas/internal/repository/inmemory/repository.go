package inmemory

import (
	"receitas/internal/entity"
	"receitas/internal/repository"
	"sort"
)

type receitaInMemory struct {
	Receitas      map[int64]entity.Receita
	ReceitaId     int64
	IngredienteId int64
}

func New(tabela map[int64]entity.Receita) repository.Interactor {
	return &receitaInMemory{
		Receitas: tabela,
	}
}

func (r *receitaInMemory) FindAll() (receitas []entity.Receita, err error) {
	for _, receita := range r.Receitas {
		receitas = append(receitas, receita)
	}
	sort.Slice(receitas, func(i, j int) bool {
		return receitas[i].Id < receitas[j].Id
	})

	return receitas, nil
}

func (r *receitaInMemory) Find(id int64) (receita entity.Receita, err error) {
	return r.Receitas[id], nil
}

func (r *receitaInMemory) Delete(id int64) error {
	delete(r.Receitas, id)
	return nil
}

func (r *receitaInMemory) Create(receita entity.Receita) (entity.Receita, error) {
	r.ReceitaId++
	id := r.ReceitaId
	receita.Id = id
	receita.Ingredientes = r.setIngredientesId(receita.Ingredientes)
	r.Receitas[id] = receita
	return receita, nil
}

func (r *receitaInMemory) setIngredientesId(ingredientes []entity.Ingrediente) []entity.Ingrediente {
	var newIngredientes []entity.Ingrediente
	for _, ingrediente := range ingredientes {
		r.IngredienteId++
		ingrediente.Id = r.IngredienteId
		newIngredientes = append(newIngredientes, ingrediente)
	}
	return newIngredientes
}

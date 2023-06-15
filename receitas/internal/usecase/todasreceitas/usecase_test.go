package todasreceitas

import (
	"receitas/internal/entity"
	"receitas/internal/repository"
	"receitas/internal/repository/inmemory"
	"reflect"
	"testing"
)

func Test_todasReceitas_BuscarReceitas(t1 *testing.T) {
	var repo = inmemory.New(map[int64]entity.Receita{})
	inserirTodasReceitasNoRepository(repo)

	type fields struct {
		repository repository.Interactor
	}
	tests := []struct {
		name   string
		fields fields
		want   []entity.Receita
	}{
		{
			name:   "Quando enviar uma receita válida, então deve retornar sucesso",
			fields: struct{ repository repository.Interactor }{repository: repo},
			want:   getTodasReceitas(),
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &todasReceitas{
				repository: tt.fields.repository,
			}
			if got := t.BuscarReceitas(); !reflect.DeepEqual(got, tt.want) {
				t1.Errorf("BuscarReceitas() = %v, want %v", got, tt.want)
			}
		})
	}
}

func inserirTodasReceitasNoRepository(repository repository.Interactor) {
	for _, receita := range getTodasReceitas() {
		repository.Create(receita)
	}
}

func getTodasReceitas() []entity.Receita {
	return []entity.Receita{
		{
			Id:     1,
			Titulo: "Bolinho de chuva",
			Preparo: "Numa tigela misture todos os ingredientes, e por último o fermento em pó, deixe descansar " +
				"por aproximadamente 5 minutos.\n Frite as colheradas em óleo não muito quente.",
			Ingredientes: []entity.Ingrediente{
				{
					Id:         1,
					Quantidade: 1,
					Descricao:  "Lata de leite condensado",
				},
				{
					Id:         2,
					Quantidade: 2,
					Descricao:  "Colher de sopa de manteiga ou margarida",
				},
				{
					Id:         3,
					Quantidade: 2,
					Descricao:  "Ovos",
				},
				{
					Id:         4,
					Quantidade: 1,
					Descricao:  "Colher de sopa de fermento em pó",
				},
				{
					Id:         5,
					Quantidade: 3,
					Descricao:  "Xicaras de chá de farinha de trigo",
				},
			},
		},
		{
			Id:     2,
			Titulo: "Bolinho de chuva",
			Preparo: "Numa tigela misture todos os ingredientes, e por último o fermento em pó, deixe descansar " +
				"por aproximadamente 5 minutos.\n Frite as colheradas em óleo não muito quente.",
			Ingredientes: []entity.Ingrediente{
				{
					Id:         6,
					Quantidade: 1,
					Descricao:  "Lata de leite condensado",
				},
				{
					Id:         7,
					Quantidade: 2,
					Descricao:  "Colher de sopa de manteiga ou margarida",
				},
				{
					Id:         8,
					Quantidade: 2,
					Descricao:  "Ovos",
				},
				{
					Id:         9,
					Quantidade: 1,
					Descricao:  "Colher de sopa de fermento em pó",
				},
				{
					Id:         10,
					Quantidade: 3,
					Descricao:  "Xicaras de chá de farinha de trigo",
				},
			},
		},
	}
}

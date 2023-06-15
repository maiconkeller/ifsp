package novareceita

import (
	"receitas/internal/entity"
	"receitas/internal/repository"
	"receitas/internal/repository/inmemory"
	"reflect"
	"testing"
)

func Test_novaReceita_Criar(t *testing.T) {
	type fields struct {
		repository repository.Interactor
	}
	type args struct {
		receita entity.Receita
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    entity.Receita
		wantErr bool
	}{
		{
			name:    "Quando enviar uma receita válida, então deve retornar sucesso",
			fields:  struct{ repository repository.Interactor }{repository: inmemory.New(map[int64]entity.Receita{})},
			args:    struct{ receita entity.Receita }{receita: getReceita()},
			want:    getReceita(),
			wantErr: false,
		},
		{
			name:    "Quando enviar uma receita sem ingredientes, então deve retornar erro",
			fields:  struct{ repository repository.Interactor }{repository: inmemory.New(map[int64]entity.Receita{})},
			args:    struct{ receita entity.Receita }{receita: getReceitaSemIngredientes()},
			want:    entity.Receita{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := &novaReceita{
				repository: tt.fields.repository,
			}
			got, err := n.Criar(tt.args.receita)
			if (err != nil) != tt.wantErr {
				t.Errorf("Criar() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Criar() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func getReceitaSemIngredientes() entity.Receita {
	receita := getReceita()
	receita.Ingredientes = nil
	return receita
}

func getReceita() entity.Receita {
	return entity.Receita{
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
	}
}

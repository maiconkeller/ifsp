package buscarreceita

import (
	"receitas/internal/entity"
	"receitas/internal/repository"
	"receitas/internal/repository/inmemory"
	"reflect"
	"testing"
)

func Test_buscarreceita_Buscar(t *testing.T) {
	type fields struct {
		repository repository.Interactor
	}
	type args struct {
		id int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    entity.Receita
		wantErr bool
	}{
		{
			name:    "Quando enviar uma ID válido, então deve retornar sucesso",
			fields:  struct{ repository repository.Interactor }{repository: inmemory.New(map[int64]entity.Receita{})},
			args:    struct{ id int64 }{id: 1},
			wantErr: false,
		},
		{
			name:    "Quando enviar uma ID zero, então deve retornar sucesso",
			fields:  struct{ repository repository.Interactor }{repository: inmemory.New(map[int64]entity.Receita{})},
			args:    struct{ id int64 }{id: 0},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &buscarreceita{
				repository: tt.fields.repository,
			}
			got, err := r.Buscar(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("Buscar() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Buscar() got = %v, want %v", got, tt.want)
			}
		})
	}
}

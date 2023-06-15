package removereceita

import (
	"receitas/internal/entity"
	"receitas/internal/repository"
	"receitas/internal/repository/inmemory"
	"testing"
)

func Test_removereceita_Excluir(t *testing.T) {
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
			r := &removereceita{
				repository: tt.fields.repository,
			}
			if err := r.Remover(tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

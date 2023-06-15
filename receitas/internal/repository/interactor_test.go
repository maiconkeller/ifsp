package repository_test

import (
	"receitas/internal/entity"
	"receitas/internal/repository"
	"receitas/internal/repository/inmemory"
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	type args struct {
		tabela map[int64]entity.Receita
	}
	tests := []struct {
		name string
		args args
		want repository.Interactor
	}{
		{
			name: "Quando chamar o new, então deve retornar sucesso",
			args: struct{ tabela map[int64]entity.Receita }{tabela: map[int64]entity.Receita{}},
			want: inmemory.New(map[int64]entity.Receita{}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := inmemory.New(tt.args.tabela); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

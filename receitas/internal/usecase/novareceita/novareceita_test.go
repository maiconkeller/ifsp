package novareceita

import (
	"receitas/internal/entity"
	"receitas/internal/repository"
	"receitas/internal/repository/inmemory"
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	type args struct {
		repository repository.Interactor
	}
	var repo = inmemory.New(map[int64]entity.Receita{})
	tests := []struct {
		name string
		args args
		want UseCase
	}{
		{
			name: "Quando chamar o new, então deve retornar sucesso",
			args: struct{ repository repository.Interactor }{repository: repo},
			want: New(repo),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.repository); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

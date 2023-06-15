package inmemory

import (
	"receitas/internal/entity"
	"reflect"
	"testing"
)

func Test_receitaInMemory_Create(t *testing.T) {
	type fields struct {
		Receitas      map[int64]entity.Receita
		ReceitaId     int64
		IngredienteId int64
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
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &receitaInMemory{
				Receitas:      tt.fields.Receitas,
				ReceitaId:     tt.fields.ReceitaId,
				IngredienteId: tt.fields.IngredienteId,
			}
			got, err := r.Create(tt.args.receita)
			if (err != nil) != tt.wantErr {
				t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Create() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_receitaInMemory_Delete(t *testing.T) {
	type fields struct {
		Receitas      map[int64]entity.Receita
		ReceitaId     int64
		IngredienteId int64
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
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &receitaInMemory{
				Receitas:      tt.fields.Receitas,
				ReceitaId:     tt.fields.ReceitaId,
				IngredienteId: tt.fields.IngredienteId,
			}
			if err := r.Delete(tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_receitaInMemory_FindAll(t *testing.T) {
	type fields struct {
		Receitas      map[int64]entity.Receita
		ReceitaId     int64
		IngredienteId int64
	}
	tests := []struct {
		name   string
		fields fields
		want   []entity.Receita
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &receitaInMemory{
				Receitas:      tt.fields.Receitas,
				ReceitaId:     tt.fields.ReceitaId,
				IngredienteId: tt.fields.IngredienteId,
			}
			if got, _ := r.FindAll(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindAll() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_receitaInMemory_setIngredientesId(t *testing.T) {
	type fields struct {
		Receitas      map[int64]entity.Receita
		ReceitaId     int64
		IngredienteId int64
	}
	type args struct {
		ingredientes []entity.Ingrediente
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []entity.Ingrediente
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &receitaInMemory{
				Receitas:      tt.fields.Receitas,
				ReceitaId:     tt.fields.ReceitaId,
				IngredienteId: tt.fields.IngredienteId,
			}
			if got := r.setIngredientesId(tt.args.ingredientes); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("setIngredientesId() = %v, want %v", got, tt.want)
			}
		})
	}
}

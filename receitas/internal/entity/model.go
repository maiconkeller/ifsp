package entity

import "time"

type Receita struct {
	Id           int64         `json:"id,omitempty"`
	Titulo       string        `json:"titulo,omitempty"`
	Preparo      string        `json:"preparo,omitempty"`
	Rendimento   string        `json:"rendimento,omitempty"`
	Ingredientes []Ingrediente `json:"ingredientes,omitempty"`
	CreatedAt    time.Time     `json:"created_at,omitempty"`
}

type Ingrediente struct {
	Id         int64  `json:"id,omitempty"`
	Quantidade int64  `json:"quantidade,omitempty"`
	Descricao  string `json:"descricao,omitempty"`
	ReceitasId int64  `json:"-"`
}

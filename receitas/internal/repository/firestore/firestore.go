package firestore

import (
	"cloud.google.com/go/firestore"
	"context"
	"receitas/internal/entity"
	"receitas/internal/repository"
)

type firestoreRepo struct {
	ctx context.Context
	cli *firestore.Client
}

func NewFirestore(ctx context.Context, cli *firestore.Client) repository.Interactor {
	return &firestoreRepo{ctx: ctx, cli: cli}
}

func (f *firestoreRepo) Create(receita entity.Receita) (entity.Receita, error) {
	//TODO implement me
	panic("implement me")
}

func (f *firestoreRepo) Delete(id int64) error {
	//TODO implement me
	panic("implement me")
}

func (f *firestoreRepo) FindAll() (receitas []entity.Receita, err error) {
	//TODO implement me
	panic("implement me")
}

func (f *firestoreRepo) Find(id int64) (receita entity.Receita, err error) {
	//TODO implement me
	panic("implement me")
}

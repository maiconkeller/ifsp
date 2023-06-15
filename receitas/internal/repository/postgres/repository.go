package postgres

import (
	"errors"
	"fmt"
	"github.com/go-pg/pg/v10"
	"github.com/labstack/gommon/log"
	"receitas/internal/entity"
	"receitas/internal/repository"
)

type receitasRepository struct {
	cli *pg.DB
}

func New(cli *pg.DB) repository.Interactor {
	return &receitasRepository{
		cli: cli,
	}
}

func (r *receitasRepository) Create(receita entity.Receita) (entity.Receita, error) {
	tx, err := r.cli.Begin()
	defer func(tx *pg.Tx) {
		err := tx.Close()
		if err != nil {
			log.Error(err)
		}
	}(tx)

	err = func() error {
		id := int64(0)
		_, err = tx.QueryOne(pg.Scan(&id), "INSERT INTO receitas (titulo, preparo, rendimento) VALUES (?,?,?) RETURNING id",
			receita.Titulo, receita.Preparo, receita.Rendimento)
		receita.Id = id

		for i := 0; i < len(receita.Ingredientes); i++ {
			receita.Ingredientes[i].ReceitasId = receita.Id
			receitaId := int64(0)
			_, err = tx.QueryOne(pg.Scan(&receitaId), "INSERT INTO ingredientes (quantidade, descricao, receitas_id) VALUES (?,?,?) RETURNING id",
				receita.Ingredientes[i].Quantidade, receita.Ingredientes[i].Descricao, receita.Ingredientes[i].ReceitasId)
			receita.Ingredientes[i].Id = receitaId
		}
		if err != nil {
			return errors.New(fmt.Sprintf("falha ao salvar a receita [%v]", err.Error()))
		}
		return nil
	}()

	if err != nil {
		_ = tx.Rollback()
		return receita, err
	}

	err = tx.Commit()
	return receita, err
}

func (r *receitasRepository) Delete(id int64) error {
	tx, err := r.cli.Begin()
	defer func(tx *pg.Tx) {
		err := tx.Close()
		if err != nil {
			log.Error(err)
		}
	}(tx)

	err = func() error {
		_, err = tx.Exec("DELETE FROM ingredientes WHERE receitas_id=?", id)
		_, err = tx.Exec("DELETE FROM receitas WHERE id=?", id)

		if err != nil {
			return errors.New(fmt.Sprintf("erro ao excluir a receita id: %v [%v]", id, err.Error()))
		}
		return nil
	}()

	if err != nil {
		_ = tx.Rollback()
		return err
	}

	return tx.Commit()
}

func (r *receitasRepository) FindAll() (receitas []entity.Receita, err error) {
	_, err = r.cli.Query(&receitas, "SELECT * FROM receitas")
	if err != nil {
		return nil, err
	}
	for i := 0; i < len(receitas); i++ {
		ingredientes, err := r.findIngredientesById(receitas[i].Id)
		if err != nil {
			return nil, err
		}
		receitas[i].Ingredientes = ingredientes
	}
	return receitas, nil
}

func (r *receitasRepository) Find(id int64) (receita entity.Receita, err error) {
	_, err = r.cli.Query(&receita, fmt.Sprintf("SELECT * FROM receitas WHERE id=%v", id))
	if err != nil {
		return entity.Receita{}, err
	}
	ingredientes, err := r.findIngredientesById(receita.Id)
	if err != nil {
		return
	}
	receita.Ingredientes = ingredientes
	return
}

func (r *receitasRepository) findIngredientesById(id int64) (ingredientes []entity.Ingrediente, err error) {
	_, err = r.cli.Query(&ingredientes, "SELECT * FROM ingredientes WHERE receitas_id=?", id)
	return
}

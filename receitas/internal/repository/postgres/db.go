package postgres

import (
	"context"
	"fmt"
	"github.com/go-pg/pg/v10"
	"time"
)

func NewDB() *pg.DB {
	opts := &pg.Options{
		Addr:     "localhost:7557",
		Database: "receitas_db",
		User:     "postgres",
		Password: "postgres",
		OnConnect: func(ctx context.Context, conn *pg.Conn) error {
			_, err := conn.Exec("SELECT 1")
			if err != nil {
				fmt.Println("postgresql is down")
				return err
			}
			return nil
		},
		ReadTimeout: 60 * time.Second,
		PoolSize:    80,
	}
	return pg.Connect(opts)
}

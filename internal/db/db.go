package db

import (
	"fmt"

	"github.com/go-pg/pg/v10"
	"github.com/rewle/service-select-participants/internal/config"
)

type DBRepository struct {
	opt  *pg.Options
	conn *pg.DB
}

func (p *DBRepository) GetConnection() *pg.DB {
	if p.conn == nil {
		p.conn = pg.Connect(p.opt)
	}
	return p.conn
}

func Init(cfg *config.Config) *DBRepository {
	fmt.Println(*cfg)
	return &DBRepository{
		&pg.Options{
			User:     cfg.DbUser,
			Password: cfg.DbPassword,
			Addr:     cfg.DbAddr,
			Database: cfg.Db,
		},
		nil,
	}
}

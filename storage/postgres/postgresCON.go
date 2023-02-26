package postgres

import (
	"fmt"
	"io/fs"
	"log"

	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
)

const (
	NotFound = "sql: no rows in result set"
)

type PostGresDB struct {
	DB *sqlx.DB
}

// Open implements fs.FS
func (*PostGresDB) Open(name string) (fs.File, error) {
	panic("unimplemented")
}

func PostGresStorageCON(cfg *viper.Viper) (*PostGresDB, error) {
	db, err := ConnectDatabase(cfg)
	if err != nil {

		return nil, err
	}
	return &PostGresDB{
		DB: db,
	}, nil
}

func ConnectDatabase(cfg *viper.Viper) (*sqlx.DB, error) {
	db, err := sqlx.Connect("postgres", fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.GetString("database.host"),
		cfg.GetString("database.port"),
		cfg.GetString("database.user"),
		cfg.GetString("database.password"),
		cfg.GetString("database.dbname"),
		cfg.GetString("database.sslmode"),
	))
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return db, nil
}

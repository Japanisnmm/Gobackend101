package middlewaresRepositories

import "github.com/jmoiron/sqlx"

type IMiddlewaresRepositories interface {
}

type middlewaresRepository struct {
	db *sqlx.DB
}
func MiddlewaresRepository(db *sqlx.DB) IMiddlewaresRepositories {
     return &middlewaresRepository {
		db: db,
	 }
}
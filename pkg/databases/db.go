package databases

import (
	"log"
	"github.com/Japanisnmm/GoBackend101/config"
	
	
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
)


func DbConnect(cfg config.IDbConfig) *sqlx.DB {
//connect
   db, err := sqlx.Connect("pgx", cfg.Url())
   if err != nil {
	log.Fatalf("connect to db failed : %v",err)
}
db.DB.SetMaxOpenConns(cfg.MaxOpenConns())
return db
}



package main

import (
	
	"os"

	"github.com/Japanisnmm/GoBackend101/config"
	"github.com/Japanisnmm/GoBackend101/modules/servers"
	"github.com/Japanisnmm/GoBackend101/pkg/databases"
)


func envPath() string {
	if len(os.Args) == 1 {
		return ".env"
	} else {
		return os.Args[1]
	}
}

func main() {
	cfg := config.LoadConfig(envPath())

	db := databases.DbConnect(cfg.Db())
	defer db.Close()

	servers.NewServer(cfg,db).Start()
	

}
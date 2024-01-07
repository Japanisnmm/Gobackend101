package config

import (
	"log"
	"math"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

func envPath() string {
	if len(os.Args) == 1 {
		return ".env"
	} else {
		return os.Args[1]
	}
}
type Iconfig interface {
	App() IAppconfig
	Db() IDbconfig
	Jwt() IJwtconfig
}




type config struct {
	app *app
	db  *db
	jwt *jwt
}

func LoadConfig(path string) Iconfig {
   envMap, err := godotenv.Read(path)
   if err != nil {
         log.Fatalf("Load dotenv failed : %v", err)
   }
	
	return &config {
		app: &app{
			host: envMap["APP_HOST"],
			port: func() int {
				p, err := strconv.Atoi(envMap["APP_PORT"])
				if err != nil {
					log.Fatalf("load port failed: %v", err)
				}
				return p
			}(),
			name : envMap["APP_NAME"],
			version: envMap["APP_VERSION"] ,
			readTimeout: func() time.Duration {
                t,err := strconv.Atoi(envMap["APP_READ_TIMEOUT"])
				  if err != nil {
					log.Fatalf("load readtimeout failed: %v", err)
				}
				 return time.Duration(int64(t) * int64(math.Pow10(9)))
			}(),
			writeTimeout: func() time.Duration {
				t,err := strconv.Atoi(envMap["APP_WRITE_TIMEOUT"])
				if err != nil {
				  log.Fatalf("load writetimeout failed: %v", err)
			  }
			   return time.Duration(int64(t) * int64(math.Pow10(9)))
		  }(),
		  bodyLimit: func() int {
			b, err := strconv.Atoi(envMap["APP_BODY_LIMIT"])
			if err != nil {
				log.Fatalf("load body limit failed: %v", err)
			}
			return b
		}(),
		fileLimit:func() int {
			b, err := strconv.Atoi(envMap["APP_FILE_LIMIT"])
			if err != nil {
				log.Fatalf("load body file limit failed: %v", err)
			}
			return b
		}(),
		gcpbucket: envMap["APP_GCP_BUCKET"],
		},
		db: &db{},
		jwt: &jwt{},
	}

} 
type IAppconfig interface {
  
}

type app struct {
	host   string
	port   int
	name    string
	version   string
	readTimeout  time.Duration
	writeTimeout time.Duration
    bodyLimit   int
	fileLimit int
	gcpbucket string

}
func (c *config)  App() IAppconfig {
	return nil

}


type IDbconfig interface {

}
type db struct {
	host string
	port int
	protocal string
	username string
	password string
	database string
	sslMode string
	maxConnection int

}
func (c *config) Db() IDbconfig {
	return nil
}


type IJwtconfig interface {
	
}

type jwt struct {
	adminKey   string
	secretKey string
	apiKey string
	accessExpireaAt int
	refreshExpireaAt int 

}
func (c *config) Jwt() IJwtconfig {
	return nil
}
package conn_pool

import (
	"database/sql"
	"fmt"
	"sync"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

type DB struct {
	conn_str string
}

var dbPass string

func init() {
	err := godotenv.Load("./vars.env")
	if err != nil {
		log.Print(".env file not found")
	}

	viper.AutomaticEnv()
	dbPass = viper.GetString("DB_PASSWORD")
	log.Print("inside init", dbPass)
}

func CreateDBConn() interface{} {
	dcs := fmt.Sprintf("root:%s@tcp(127.0.0.1:3306)/sys?parseTime=true&timeout=5s", dbPass)
	// log.Print(dcs)

	db, err := sql.Open("mysql", dcs)
	if err != nil {
		log.Printf("Error opening the connection %v", err)
	}

	db.SetMaxIdleConns(1)
	db.SetMaxOpenConns(1)

	return db
}

func GetDBPool() *sync.Pool {
	p := &sync.Pool{
		New: CreateDBConn,
	}

	for i := 0; i <= 10; i += 1 {
		p.Put(p.New())
	}
	return p
}

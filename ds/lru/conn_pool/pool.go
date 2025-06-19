package conn_pool

import (
	"database/sql"
	"sync"

	_ "github.com/go-sql-driver/mysql"
	"github.com/rs/zerolog/log"
)

type DB struct {
	conn_str string
}

func CreateDBConn() interface{} {
	dcs := ""

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

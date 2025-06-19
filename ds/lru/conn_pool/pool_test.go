package conn_pool

import (
	"database/sql"
	"log"
	"sync"
	"testing"
)

func BenchmarkWithPooling(b *testing.B) {
	var wg sync.WaitGroup
	var mu sync.Mutex

	var rc int
	pool := GetDBPool()
	for i := 0; i < b.N; i += 1 {
		wg.Add(1)
		log.Println("at i = ", i)
		go func() {
			defer wg.Done()
			dbconn := pool.Get().(*sql.DB)
			rows, err := dbconn.Query("SELECT * from customer")
			if err != nil {
				log.Println("rc = ", rc)
				log.Fatalf("Error retreiving from database! %v", err)
			}

			defer rows.Close()
			var customer_id int
			var name string
			var visited_on string
			var amount int

			for rows.Next() {
				mu.Lock()
				rc += 1
				mu.Unlock()

				if err := rows.Scan(&customer_id, &name, &visited_on, &amount); err != nil {
					log.Println("rc = ", rc)
					log.Fatalf("error retriving value from row %v", err)
				}
				log.Print(customer_id, " ", name)
			}

			pool.Put(dbconn)
		}()
	}

	wg.Wait()

}

func BenchmarkWithoutPooling(b *testing.B) {
	var wg sync.WaitGroup
	var mu sync.Mutex
	var rc int

	for i := 0; i < b.N; i += 1 {
		wg.Add(1)
		log.Println("at i = ", i)
		go func() {
			defer wg.Done()
			dbconn := CreateDBConn().(*sql.DB)
			defer dbconn.Close()
			rows, err := dbconn.Query("SELECT * from customer")
			if err != nil {
				log.Println("rc = ", rc)
				log.Fatalf("Error retreiving from database! %v", err)
				return
			}

			defer rows.Close()
			var customer_id int
			var name string
			var visited_on string
			var amount int

			for rows.Next() {
				mu.Lock()
				rc += 1
				mu.Unlock()

				if err := rows.Scan(&customer_id, &name, &visited_on, &amount); err != nil {
					log.Println("rc = ", rc)
					log.Fatalf("error retriving value from row %v", err)
				}
				log.Print(customer_id, " ", name)
			}

		}()
	}

	wg.Wait()

}

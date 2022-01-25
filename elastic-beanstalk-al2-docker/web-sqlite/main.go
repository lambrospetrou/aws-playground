package main

import (
	"database/sql"
	"fmt"
	"html"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"

	_ "modernc.org/sqlite"
)

func createTable(db *sql.DB) {
	sqlStmt := `
	create table if not exists users (name TEXT, age INTEGER);
	--delete from users;
	`
	_, err := db.Exec(sqlStmt)
	if err != nil {
		log.Fatalf("%q: %s\n", err, sqlStmt)
	}
}

func transactionInserts(db *sql.DB, num int) {
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	stmt, err := tx.Prepare("INSERT INTO users (name, age) VALUES (?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	for i := 0; i < num; i++ {
		var suffix string
		if rand.Float64() >= 0.6 {
			suffix = "Lambros"
		} else {
			suffix = fmt.Sprintf("Lambros-%03d", i)
		}
		_, err = stmt.Exec(suffix, i)
		if err != nil {
			log.Fatal(err)
		}
	}
	tx.Commit()
	log.Printf("finished inserting %d users", num)
}

func queryAll(db *sql.DB) {
	rows, err := db.Query("SELECT * FROM users WHERE name LIKE ?", "Lambros-%")
	if err != nil {
		log.Fatal(err)
	}
	n := 0
	for rows.Next() {
		n++
	}
	log.Printf("finished fetching %d rows", n)

	rows, err = db.Query("SELECT count(*) FROM users")
	if err != nil {
		log.Fatal(err)
	}
	rows.Next()
	cnt := 0
	if err = rows.Scan(&cnt); err != nil {
		log.Fatal(err)
	}
	log.Println("all rows:", cnt)
}

func timeIt(prefix string, f func()) time.Duration {
	start := time.Now()
	f()
	timeElapsed := time.Since(start)
	log.Println("Time elapsed in", prefix, "is:", timeElapsed)
	return timeElapsed
}

func handleRequest(db *sql.DB) time.Duration {
	var err error
	var numInserts = 100
	numInsertsStr := os.Getenv("NUM_INSERTS")
	if numInsertsStr != "" {
		numInserts, err = strconv.Atoi(numInsertsStr)
		if err != nil {
			log.Fatal(err)
		}
	}

	totalDuration := timeIt("main()", func() {
		// createTable(db)
		timeIt("inserts", func() {
			transactionInserts(db, numInserts)
		})
		timeIt("query", func() {
			queryAll(db)
		})
	})
	return totalDuration
}

func main() {
	dbPath := os.Getenv("SQLITE_PATH")
	if dbPath == "" {
		dbPath = "./users.db"
	}

	// os.Remove(dbPath)
	db, err := sql.Open("sqlite", dbPath)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	createTable(db)

	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "5000"
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		totalDuration := handleRequest(db)
		fmt.Fprintf(w, "Service SQLite Path - queries done - ms[%d]!, %q", totalDuration.Milliseconds(), html.EscapeString(r.URL.Path))
	})

	log.Println("Web Service SQLite starts listening at :" + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

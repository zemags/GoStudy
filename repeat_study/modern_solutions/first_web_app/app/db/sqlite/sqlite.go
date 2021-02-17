package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	// "github.com/joho/godotenv"
	"github.com/joho/godotenv"
	_ "github.com/mattn/go-sqlite3"

	"github.com/zemags/GoStudy/repeat_study/modern_solutions/first_web_app/app/config"
)

type planet struct {
	id              int
	planetType      string
	planetShortName string
	zone            int
	age             int
}

// init is invoked before main()
func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	conf := config.NewConfig()
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	dbName := fmt.Sprintf("%s/%s.db", dir, conf.DBParams.DBName)
	fmt.Println(dbName, conf.DBParams.DBName)
	db, err := sql.Open("sqlite3", dbName)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query("select planet_short_name from planets where id = $1", 1)
	if err != nil {
		log.Fatal(err)
	}
	planetes := []planet{}
	for rows.Next() {
		p := planet{}
		if err := rows.Scan(&p.planetShortName); err != nil {
			log.Fatal(err)
		}
		planetes = append(planetes, p)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println(planetes)
}

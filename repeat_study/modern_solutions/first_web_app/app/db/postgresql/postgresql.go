package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
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
	postgresqlConn := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", conf.DBParams.User, conf.DBParams.Password, conf.DBParams.DBName)

	db, err := sql.Open("postgres", postgresqlConn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query("select * from planets where id = $1", 2)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	planets := []planet{}
	for rows.Next() {
		p := planet{}
		err := rows.Scan(&p.id, &p.planetType, &p.planetShortName, &p.zone, &p.age)
		if err != nil {
			log.Println(err)
			continue
		}
		planets = append(planets, p)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println(planets)

	// get on row
	// row := db.QueryRow("select * from planets where age > $1", 4500000000)

	// insert rows
	_, errInsert := db.Exec("insert into planets (planet_type, planet_short_name, zone, age) values ('comet1', 'halley1', 2, 2261)")
	if errInsert != nil {
		log.Fatal(errInsert)
	}

	// update row
	newAge := 4600000004
	resultUpdate, err := db.Exec("update planets set age = $1 where id = $2", newAge, 2)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resultUpdate.RowsAffected())

	// with prepare can user many times
	stmt, err := db.Prepare("select * from planets where planet_type = $1")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	rows, err = stmt.Query(1)
	if err != nil {
		log.Fatal(err)
	}
	rows, err = stmt.Query(2)
	if err != nil {
		log.Fatal(err)
	}

	//transaction
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	defer tx.Rollback()
	stmt1, err := tx.Prepare("select * from planets where planet_type = $1")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt1.Close()
	rows, err = stmt1.Query(1)
	if err != nil {
		log.Fatal(err)
	}
	if err := tx.Commit(); err != nil {
		log.Fatal(err)
	}
}

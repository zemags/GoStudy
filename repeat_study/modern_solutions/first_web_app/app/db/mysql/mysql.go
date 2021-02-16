package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
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
	db, err := sql.Open("mysql", conf.DBParams.MysqlDBName)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query("select * from universe.planets where id=?", 1)
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
	// row := db.QueryRow("select * from universe.planets where age > ?", 4500000000)

	// insert rows
	resultInsert, err := db.Exec("insert into planets (planet_type, planet_short_name, zone, age) values ('gas', 'saturn', 2, 4600000003)")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resultInsert.LastInsertId())

	// update row
	newAge := 4600000004
	resultUpdate, err := db.Exec("update universe.planets set age=? where id=?", newAge, 5)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resultUpdate.LastInsertId())
	fmt.Println(resultUpdate.RowsAffected())

}

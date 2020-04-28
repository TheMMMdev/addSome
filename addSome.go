package main

import (
	"bufio"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"os"
)

const (
	host     = "<your host ip>"
	port     = 5432
	user     = "<postgres username>"
	password = "<postgres password>"
	dbname   = "<name of database to compare"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("No file specified, exiting...")
		os.Exit(2)
	}
	targetFile := os.Args[1]

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("connected to psql")

	file, err := os.Open(targetFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {

			insertStatement := `
			INSERT INTO subdomains (name) VALUES ($1) on conflict (name) do nothing`

			id, err := db.Exec(insertStatement, scanner.Text())
			rows, err := id.RowsAffected()
			if err != nil {
				fmt.Println(err)
			}
			if err != nil {
				panic(err)
			}
			if rows != 0 {
				fmt.Printf("%s added!\n", scanner.Text())
			}

			if err := scanner.Err(); err != nil {
				panic(err)
			}

	}

}

package main

import (
	"database/sql"
	"fmt"
	"log"

	// not used
	"github.com/deeper-x/themock/mydb"
	_ "github.com/lib/pq"
)

func main() {
	var fkShip string
	dbc, err := sql.Open("postgres", "user=demo dbname=demo password=demo sslmode=disable")

	if err != nil {
		log.Println(err)
	}

	fmt.Println(dbc)

	conn := mydb.NewConnection(dbc)

	rows := conn.DemoQuery("28")

	for rows.Next() {
		err = rows.Scan(
			&fkShip,
		)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(fkShip)
	}

}

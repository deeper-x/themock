package mydb

import (
	"database/sql"
	"log"
)

// LiveData live data container
type LiveData interface {
	DemoQuery(idPortinformer string) *sql.Rows
}

type connection struct {
	db *sql.DB
}

// NewConnection is connection startpoint
func NewConnection(db *sql.DB) LiveData {
	return &connection{db: db}
}

func (c connection) DemoQuery(idPortinformer string) *sql.Rows {
	var fkShip int
	query := "SELECT fk_ship from control_unit_data WHERE fk_portinformer = $1"

	rows, err := c.db.Query(query, idPortinformer)

	if err != nil {
		log.Println(err)
	}

	rows.Scan(
		&fkShip,
	)

	if err != nil {
		log.Println(err)
	}

	return rows
}

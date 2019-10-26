package mydb

import (
	"log"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestDemoQuery(t *testing.T) {
	db, mock, err := sqlmock.New()

	if err != nil {
		log.Println(err)
	}

	defer db.Close()

	expectedRows := sqlmock.NewRows([]string{"fk_ship"})
	mock.ExpectQuery("SELECT fk_ship").
		WithArgs("28").
		WillReturnRows(expectedRows)

	mocked := NewConnection(db)
	mocked.DemoQuery("28")

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("%v", err)
	}
}

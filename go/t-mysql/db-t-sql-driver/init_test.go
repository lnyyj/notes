package dbtsqldriver

import (
	"database/sql"
	"fmt"
	"log"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

var (
	db  *sql.DB
	err error
)

func init() {
	if db, err = sql.Open("mysql", "root:1234qwer@tcp(localhost:3306)/admin"); err != nil {
		panic(err.Error())
	}
}
func Test_db_query(t *testing.T) {
	rows, err := db.Query("select * from orders limit 1")
	if err != nil {
		fmt.Printf("---->[%+v]\r\n", err)
		return
	}
	fmt.Printf("----->query[%+v]\r\n", rows)
}

func Test_db_query_row(t *testing.T) {
	var (
		id   int
		name string
	)
	rows, err := db.Query("select id, name from customers where id = ?", 15889337792)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&id, &name)
		if err != nil {
			log.Fatal(err)
		}
		t.Log(id, name)
	}
	err = rows.Err()
	if err != nil {
		t.Fatal(err)
	}
}

func Test_db_query_stmt(t *testing.T) {
	stmt, err := db.Prepare("select id, name from customers where id = ?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	rows, err := stmt.Query(15889337792)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var (
		id   int
		name string
	)
	for rows.Next() {
		err := rows.Scan(&id, &name)
		if err != nil {
			log.Fatal(err)
		}
		t.Log("---->", id, name)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}
}

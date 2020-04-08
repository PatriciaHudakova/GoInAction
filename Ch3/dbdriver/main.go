package main

import (
	"database/sql"
	_ "github.com/goinaction/code/chapter3/dbdriver/postgres" //blank identifier to prevent compile error
)

/*func init() {
	sql.Register("postgres", new(PostgresDriver))        //driver initialised in sql package
} */

func main() {
	sql.Open("postgres", "mydb")
}

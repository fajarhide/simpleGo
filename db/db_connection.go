package db
 
import (
	"database/sql"
	"fmt"
	_"github.com/go-sql-driver/mysql"
)
 
/*Create mysql connection*/
func CreateCon() *sql.DB {
	db, err := sql.Open("mysql", "user:userpass@tcp(127.0.0.1:3306)/golang")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("db is connected")
	}
	//defer db.Close()
	// make sure connection is available
	err = db.Ping()
	fmt.Println(err)
	if err != nil {
		fmt.Println("db is not connected")
		fmt.Println(err.Error())
	}
	return db
}
 
/*end mysql connection*/
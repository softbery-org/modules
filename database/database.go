// Softbery Libraries (c) by Paweł Tobis
//
// Softbery Libraries is licensed under a
// Creative Commons Attribution-ShareAlike 4.0 International License.
//
// You should have received a copy of the license along with this
// work. If not, see <http://creativecommons.org/licenses/by-sa/4.0/>.

package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func Connect() bool {
	result := false

	var dberror error
	db, dberror = sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/golang_modules")

	// sql.Open("type", "username:password@net(address:port)/dbname")
	// sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/golang_modules")
	// db, dberror := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/golang_modules")

	pingError := db.Ping()

	if dberror != nil {
		log.Fatal(dberror)
		log.Fatal(pingError)
		result = false
	} else {
		defer db.Close()
		fmt.Println("Connected to database!!!")
		result = true
	}
	return result
}

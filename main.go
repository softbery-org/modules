// Softbery Libraries (c) by Paweł Tobis
//
// Softbery Libraries is licensed under a
// Creative Commons Attribution-ShareAlike 4.0 International License.
//
// You should have received a copy of the license along with this
// work. If not, see <http://creativecommons.org/licenses/by-sa/4.0/>.

package main

import (
	"database/sql"
	"fmt"

	"softbery.org/database"
)

var db *sql.DB

func main() {
	conn := database.Connect()
	fmt.Println(conn)
}

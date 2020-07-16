package dependency

import (
	"fmt"
)


// DatabaseMySQL is an concret implementation of db connection
type DatabaseMySQL struct {}


// Connect is the method for connect to db
func (db *DatabaseMySQL) Connect() {
	fmt.Println("Connected to MySQL")
}
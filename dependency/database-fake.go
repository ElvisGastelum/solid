package dependency

import (
	"fmt"
)


// DatabaseFake is an concret implementation of db connection
type DatabaseFake struct {}

// Connect is the method for connect to db
func (db *DatabaseFake) Connect() {
	fmt.Println("Connected to Mock")
}
package dependency

import (
	"fmt"
)


// DatabaseFirebase is an concret implementation of db connection
type DatabaseFirebase struct {}

// Connect is the method for connect to db
func (db *DatabaseFirebase) Connect() {
	fmt.Println("Connected to Firebase")
}
package dependency

import (
	"fmt"
)


// DatabaseMongoDB is an concret implementation of db connection
type DatabaseMongoDB struct {}

// Connect is the method for connect to db
func (db *DatabaseMongoDB) Connect() {
	fmt.Println("Connected to MongoDB")
}
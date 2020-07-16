package dependency


// Dependency this have the connection to the database
type Dependency struct{}

 
// Start method to connect
func (d *Dependency) Start(db Database){
	db.Connect()
}


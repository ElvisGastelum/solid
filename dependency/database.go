package dependency

// Database is an interface to connect to whatever db
type Database interface{
	Connect()
}



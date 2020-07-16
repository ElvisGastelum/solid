package main

import "github.com/elvisgastelum/solid/dependency"

func main() {
	db := dependency.DatabaseFake{}

	dp := dependency.Dependency{}

	dp.Start(&db)



}
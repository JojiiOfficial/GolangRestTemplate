package main

import (
	dbhelper "github.com/JojiiOfficial/GoDBHelper"
)

func getInitSQL() dbhelper.QueryChain {
	return dbhelper.QueryChain{
		Name:    "initChain",
		Order:   0,
		Queries: dbhelper.CreateInitVersionSQL(),
	}
}
func isConnectedToDB() error {
	return nil
}

package database

import (
	"log"
	"upper.io/db.v3"
	"upper.io/db.v3/mysql"
)

var databaseConfig = mysql.ConnectionURL{
	User:     `root`,
	Database: `scrumpoker`,
	Host:     `localhost:6603`,
	Password: `scrumUmad9001`,
}

func NewCollection(name string) (db.Collection, error) {
	conn, err := getDatabaseConnection()
	if nil != err {
		return nil, err
	}

	return conn.Collection(name), nil
}

var databaseConnection db.Database

func getDatabaseConnection() (db.Database, error) {
	if nil == databaseConnection {
		var err error
		databaseConnection, err = mysql.Open(databaseConfig)
		if err != nil {
			log.Fatalf("db.Open(): %q\n", err)
			return nil, err
		}

		//defer databaseConnection.Close()
	}

	return databaseConnection, nil
}

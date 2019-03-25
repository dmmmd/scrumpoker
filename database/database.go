package database

import (
	"log"
	"upper.io/db.v3"
	"upper.io/db.v3/mysql"
)

var databaseConfig = mysql.ConnectionURL{
	User:     `scrumpoker`,
	Database: `scrumpoker`,
	Host:     `mysql:3306`,
	Password: `hYnRDFfWGdPCCG8BcKpZUvxWz6YaM3`,
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

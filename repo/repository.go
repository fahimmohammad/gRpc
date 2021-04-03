package repo

import (
	"fmt"

	"github.com/fahimsGit/gRpctest/config"
	"github.com/globalsign/mgo"
)

type DBSession struct {
	DbSession *mgo.Session
	DbName    string
	TableName string
}

func NewDB() *DBSession {
	session, err := mgo.Dial("localhost:27017")
	if err != nil {
		fmt.Println("Could not connct")
		return &DBSession{
			DbSession: &mgo.Session{},
			DbName:    "",
			TableName: "",
		}
	}
	configuration := config.NewConfig()
	return &DBSession{
		DbSession: session,
		DbName:    configuration.DbName,
		TableName: configuration.TableName,
	}

}

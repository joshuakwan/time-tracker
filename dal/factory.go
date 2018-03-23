package db

import (
	"github.com/joshuakwan/time-tracker/config"
	mgo "gopkg.in/mgo.v2"
)

// DataSession defines the common interface for data access
type DataSession interface {
	CreateObject(tableName string, object interface{}) error
	RemoveObject(tableName string, query string) error
	UpdateObject(tableName string, selector interface{}, update interface{}) error
	FindObjects(tableName string, query string, result interface{}) error
	GetAllObjects(colName string, result interface{}) error
}

var dbSession *mongoDBSession

// GetSession returns the concrete data access session
func GetSession() DataSession {
	databaseType := config.GetDatabaseType()
	if databaseType == "mongo" {
		if dbSession == nil {
			mongoSession, err := mgo.Dial(config.GetDatabaseURL())
			if err != nil {
				panic(err)
			}
			mongoSession.SetMode(mgo.Monotonic, true)
			mongoSession.SetSafe(&mgo.Safe{})
			dbSession = &mongoDBSession{session: mongoSession}
		}
		return dbSession
	}

	return nil
}

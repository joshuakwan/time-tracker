package db

import (
	"strings"

	"gopkg.in/mgo.v2/bson"

	"github.com/joshuakwan/time-tracker/config"
	"gopkg.in/mgo.v2"
)

type mongoDBSession struct {
	session *mgo.Session
}

// CreateObject creates a new mongo object
func (db *mongoDBSession) CreateObject(colName string, object interface{}) error {
	c := db.session.DB(config.GetDatabaseName()).C(colName)
	return c.Insert(object)
}

// RemoveObject removes a DB object
func (db *mongoDBSession) RemoveObject(colName string, query string) error {
	parts := strings.Split(query, ":")

	col := db.session.DB(config.GetDatabaseName()).C(colName)
	err := col.Remove(bson.M{parts[0]: parts[1]})
	return err
}

// UpdateObject updates a DB object
func (db *mongoDBSession) UpdateObject(colName string, selector interface{}, update interface{}) error {
	col := db.session.DB(config.GetDatabaseName()).C(colName)
	err := col.Update(selector, update)
	return err
}

// FindObjects finds a DB object
func (db *mongoDBSession) FindObjects(colName string, query string, result interface{}) error {
	parts := strings.Split(query, ":")

	col := db.session.DB(config.GetDatabaseName()).C(colName)
	err := col.Find(bson.M{parts[0]: parts[1]}).All(result)
	return err
}

// GetAllObjects returns all DB objects from a collection
func (db *mongoDBSession) GetAllObjects(colName string, result interface{}) error {
	col := db.session.DB(config.GetDatabaseName()).C(colName)
	err := col.Find(nil).All(result)
	return err
}

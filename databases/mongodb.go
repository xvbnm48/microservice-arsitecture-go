package databases

import "gopkg.in/mgo.v2"

type MongoDB struct {
	MgDbSession  *mgo.Session
	Databasename string
}

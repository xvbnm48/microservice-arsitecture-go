package databases

import (
	"time"

	"github.com/raycad/go-microservices/tree/master/src/movie-microservice/common"
	logging "github.com/sirupsen/logrus"
	"gopkg.in/mgo.v2"
)

type MongoDB struct {
	MgDbSession  *mgo.Session
	Databasename string
}

func (db *MongoDB) Init() error {
	db.Databasename = common.Config.MgDbName

	// dial info holds options for establishing a session with a Mongo db cluster
	dialInfo := &mgo.DialInfo{
		Addrs:    []string{common.Config.MgAddrs},
		Timeout:  60 * time.Second,
		Database: db.Databasename,
		Username: common.Config.MgDbUserName,
		Password: common.Config.MgDbPassword,
	}

	// create a new session which maintains a pool of socket connections
	// to the MongoDB cluster
	var err error
	db.MgDbSession, err = mgo.DialWithInfo(dialInfo)

	if err != nil {
		logging.Debug("cant cannot connect mongo, go error: ", err)
		return err
	}

	return err
}

func (db *MongoDB) Close() {
	if db.MgDbSession != nil {
		db.MgDbSession.Close()
	}
}

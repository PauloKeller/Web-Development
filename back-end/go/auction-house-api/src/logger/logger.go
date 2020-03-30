package logger

import (
	"fmt"
	"log"
	"time"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// MongoLogger global instance
var MongoLogger *log.Logger

// MongoWriter receives a mongo session to write files in mongo db
type MongoWriter struct {
	sess *mgo.Session
}

func init() {
	var err error
	MongoLogger, err = NewMongoLogger()

	if err != nil {
		panic(err)
	}
}

// NewMongoLogger creates a new logger mongo
func NewMongoLogger() (*log.Logger, error) {

	const (
		hosts      = "mongodb:27017"
		database   = "logger"
		collection = "logs"
	)

	info := &mgo.DialInfo{
		Addrs:    []string{hosts},
		Timeout:  60 * time.Second,
		Database: database,
	}

	sess, err := mgo.DialWithInfo(info)
	if err != nil {
		panic(err)
	}

	if err = sess.Ping(); err != nil {
		panic(err)
	}

	fmt.Println("Logger database initialized.")

	mw := &MongoWriter{sess}
	logger := log.New(mw, "API LOGGER-", log.LstdFlags|log.Lshortfile)

	return logger, nil
}

func (mw *MongoWriter) Write(p []byte) (n int, err error) {
	c := mw.sess.DB("logger").C("logs")
	err = c.Insert(bson.M{
		"created": time.Now(),
		"msg":     string(p),
	})
	if err != nil {
		return
	}
	return len(p), nil
}

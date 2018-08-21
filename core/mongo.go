package core

import (
	"github.com/pkg/errors"
	"gopkg.in/mgo.v2"
	"strings"
	"time"
)

type MongoConfig struct {
	Hosts    []string `env:"MONGO_HOSTS,required" envSeparator:","`
	AuthDb   string   `env:"MONGO_AUTH_DB" envDefault:"admin"`
	Database string   `env:"MONGO_DATABASE"`
	Username string   `env:"MONGO_USER,required"`
	Password string   `env:"MONGO_PASS,required"`
}

func (mongoConfig *MongoConfig) CreateSession() (DB, error) {
	db, err := mgo.DialWithInfo(mongoConfig.GetDialInfo())

	if err != nil {
		err = errors.Wrapf(err, "cannot dial mongo at (%s)", strings.Join(mongoConfig.Hosts, ","))
		return nil, err
	}

	session := &MongoSession{db}

	if err := session.Ping(); err != nil {
		return session, err
	}

	return session, nil
}

func (mongoConfig *MongoConfig) GetDialInfo() *mgo.DialInfo {
	return &mgo.DialInfo{
		Addrs:    mongoConfig.Hosts,
		Timeout:  10 * time.Second,
		Database: mongoConfig.AuthDb,
		Username: mongoConfig.Username,
		Password: mongoConfig.Password,
	}
}

type DB interface {
	DB(name string) DataLayer
	Copy() DB
	Close()
}

type MongoSession struct {
	*mgo.Session
}

// DB wraps *mgo.DB to returns a DataLayer interface instead of *mgo.Database.
func (s MongoSession) DB(name string) DataLayer {
	return &MongoDataLayer{Database: s.Session.DB(name)}
}

func (s MongoSession) Copy() DB {
	return MongoSession{s.Session.Copy()}
}

func (s MongoSession) Close() {
	s.Session.Close()
}

// DataLayer is an interface to access to the database struct.
type DataLayer interface {
	C(name string) Collection
}

// MongoDataLayer wraps a mgo.Database to embed methods in models.
type MongoDataLayer struct {
	*mgo.Database
}

// C wraps *mgo.DB to return a DataLayer interface instead of *mgo.Database.
func (mongo MongoDataLayer) C(name string) Collection {
	return &MongoCollection{Collection: mongo.Database.C(name)}
}

type MongoCollection struct {
	*mgo.Collection
}

// Collection is an interface to access to the collection struct.
type Collection interface {
	Find(query interface{}) *mgo.Query
	Insert(docs ...interface{}) error
	Update(selector interface{}, update interface{}) error
}

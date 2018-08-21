package core

import "gopkg.in/mgo.v2"

type MockMongo struct {
	Assertable AssertableMongo
	Mockable   MockableMongo
}

func (m *MockMongo) DB(name string) DataLayer {
	m.Assertable.Name = name
	return m.Mockable.Datalayer
}

func (m *MockMongo) Copy() DB {
	return m
}

func (m *MockMongo) Close() {
	m.Assertable.IsClosed = true
}

type AssertableMongo struct {
	IsClosed bool
	Name     string
}

type MockableMongo struct {
	Datalayer DataLayer
}

type MockDatalayer struct {
	AssertableName     string
	MockableCollection Collection
}

func (d *MockDatalayer) C(name string) Collection {
	d.AssertableName = name
	return MongoCollection{}
}

type MockCollection struct {
	Assertable AssertableCollection
	Mockable   MockableCollection
}

func (c *MockCollection) Find(query interface{}) *mgo.Query {
	c.Assertable.Query = query
	return c.Mockable.Query
}

func (c *MockCollection) Insert(docs ...interface{}) error {
	c.Assertable.Docs = docs
	return c.Mockable.Err
}

func (c *MockCollection) Update(selector interface{}, update interface{}) error {
	c.Assertable.Selector = selector
	c.Assertable.Update = update
	return c.Mockable.Err
}

type AssertableCollection struct {
	Query    interface{}
	Docs     []interface{}
	Selector interface{}
	Update   interface{}
}

type MockableCollection struct {
	Query *mgo.Query
	Err   error
}

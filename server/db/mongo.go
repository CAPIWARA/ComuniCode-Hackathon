package db

import (
	"errors"
	"log"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var Session *mgo.Session
var Database = "comunicode"

type MongoCollection struct {
	*mgo.Collection
}

func (conn *MongoCollection) FindById(id string) (interface{}, error) {
	var data map[string]interface{}

	if err := conn.Find(bson.M{"_id": bson.ObjectIdHex(id)}).One(&data); err != nil {
		return nil, err
	}

	data["id"] = data["_id"].(bson.ObjectId).Hex()
	return data, nil
}

func (conn *MongoCollection) FindByQuery(query string, value string) (interface{}, error) {
	var data map[string]interface{}

	if err := conn.Find(bson.M{query: value}).One(&data); err != nil {
		return nil, err
	}

	data["id"] = data["_id"].(bson.ObjectId).Hex()
	return data, nil
}

func (conn *MongoCollection) Save(obj interface{}) error {
	var result interface{}

	payload, err := InterfaceToMap(obj)
	if err != nil {
		return err
	}

	id := bson.NewObjectId()
	(*payload)["_id"] = id
	delete(*payload, "id")

	err = conn.Insert(payload)
	if err != nil {
		if mgo.IsDup(err) {
			return errors.New("record already exists!")
		}
		return err
	}

	(*payload)["id"] = id.Hex()
	err = MapToInterface(payload, &obj)
	if err != nil {
		return err
	}

	return nil
}

func MongoRepoBuilder(repoDefinition RepositoryDef) Repository {
	mongoCollection, err := prepareDB(Session, Database, repoDefinition.GetName())
	if err != nil {
		log.Printf("error in mongo builder: %v", err)
		return nil
	}
	return &MongoCollection{
		mongoCollection,
	}
}

func prepareDB(session *mgo.Session, db, dbcollection string) (*mgo.Collection, error) {
	if session == nil {
		return nil, errors.New("session is nil")
	}
	return session.DB(db).C(dbcollection), nil
}

func NewSession() error {
	session, err := mgo.Dial("mongodb://cardosomarcos:cardosomarcos10@ds153890.mlab.com:53890/comunicode")

	if err != nil {
		log.Printf("db error: %v", err)
		return err
	}
	session.SetMode(mgo.Monotonic, true)
	Session = session
	return nil
}

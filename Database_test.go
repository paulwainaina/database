package database

import (
	"fmt"
	"testing"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var database = Database{DbName: "example", DbServer: "mongodb://localhost:27017"}

func TestConnect(t *testing.T) {
	if !database.Connect() {
		t.Errorf("Error connecting %q ", database.Err)
	}
}

func TestPing(t *testing.T) {
	database.Connect()
	if !database.Ping() {
		t.Errorf("Error pinging %q", database.Err)
	}
}

func TestRequiredString(t *testing.T) {
	if len(database.DbName) <= 0 {
		t.Errorf("Error setting Database Name")
	}
	if len(database.DbServer) <= 0 {
		t.Errorf("Error setting Database server")
	}
}

func TestDisconnect(t *testing.T) {
	database.Connect()
	database.Disconnect()
	if database.Err != nil {
		t.Errorf("Error closing connection %q", database.Err)
	}
}

type Address struct {
	City string
}
type Person struct {
	Name    string  `bson:"name,omitempty"`
	Address Address `bson:"inline"`
}

func TestQueryOne(t *testing.T) {
	database.Connect()
	filter := bson.M{}
	result := database.QueryOne(database.DbName, "users", filter)
	if database.Err != nil && database.Err != mongo.ErrNoDocuments {
		t.Errorf("Error closing connection %q", database.Err)
	}
	var person Person
	bsonBytes, err := bson.Marshal(result)
	if err != nil {
		t.Errorf("Error closing connection %q", err)
	}
	bson.Unmarshal(bsonBytes, &person)
	fmt.Println(person)
}

func TestInsertOne(t *testing.T) {
	database.Connect()
	person := Person{Name: "test", Address: Address{City: "NewYork"}}
	database.InsertOne(database.DbName, "users", person)
	if database.Err != nil {
		t.Errorf("Error closing connection %q", database.Err)
	}
	filter := bson.M{"name": "test"}
	database.QueryOne(database.DbName, "users", filter)
	if database.Err != nil {
		t.Errorf("Error closing connection %q", database.Err)
	}
}

func TestInsertManyQueryMany(t *testing.T) {
	database.Connect()
	persons := []interface{}{
				Person{Name: "test", Address: Address{City: "NewYork"}},
				Person{Name: "test1", Address: Address{City: "Nairobi"}},
			}
	database.InsertMany(database.DbName,"users",persons)
	if database.Err != nil {
		t.Errorf("Error closing connection %q", database.Err)
	}
	filter := bson.M{"$or":[]bson.M{{"name":"test"},{"name":"test1"}}}
	database.QueryMany(database.DbName, "users", filter)
	if database.Err != nil {
		t.Errorf("Error closing connection %q", database.Err)
	}
}

func TestUpdateOne(t *testing.T){
	database.Connect()
	person := Person{Name: "test", Address: Address{City: "NewYork"}}
	database.InsertOne(database.DbName, "users", person)
	if database.Err != nil {
		t.Errorf("Error closing connection %q", database.Err)
	}
	filter:=bson.M{"name":"test"}
	update:=bson.M{"$set":bson.M{"name":"test1"}}
	var options *options.UpdateOptions
	database.UpdateOne(database.DbName,"users",filter,update,options)
	if database.Err != nil {
		t.Errorf("Error closing connection %q", database.Err)
	}
}

func TestUpdateMany(t *testing.T){
	database.Connect()
	person := Person{Name: "test", Address: Address{City: "NewYork"}}
	database.InsertOne(database.DbName, "users", person)
	if database.Err != nil {
		t.Errorf("Error closing connection %q", database.Err)
	}
	filter:=bson.M{}
	update:=bson.M{"$set":bson.M{"name":"test2"}}
	var options *options.UpdateOptions
	database.UpdateMany(database.DbName,"users",filter,update,options)
	if database.Err != nil {
		t.Errorf("Error closing connection %q", database.Err)
	}
}

func TestDeleteOne(t *testing.T){
	database.Connect()
	filter := bson.M{"$or":[]bson.M{{"name":"test"},{"name":"test1"}}}
	var options *options.DeleteOptions
	database.DeleteOne(database.DbName,"users",filter,options)
	if database.Err != nil {
		t.Errorf("Error closing connection %q", database.Err)
	}
}

func TestDeleteMany(t *testing.T){
	database.Connect()
	filter := bson.M{"$or":[]bson.M{{"name":"test"},{"name":"test1"}}}
	var options *options.DeleteOptions
	database.DeleteMany(database.DbName,"users",filter,options)
	if database.Err != nil {
		t.Errorf("Error closing connection %q", database.Err)
	}
}

func TestFindDistinct(t *testing.T){
	database.Connect()
	var filter string="name"
	options:=  options.Distinct().SetMaxTime(2 * time.Second)
	result:=database.FindDistinct(database.DbName,"users",filter,options)
	if database.Err != nil {
		t.Errorf("Error closing connection %q", database.Err)
	}
	fmt.Println(result)
}
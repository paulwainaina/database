package database

import (
	"context"
	"time"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type Database struct {
	DbServer string
	DbName   string
	Ctx      context.Context
	Cancel   context.CancelFunc
	Client   *mongo.Client
	Err      error
}

func (db *Database) Connect() bool {

	db.Client, db.Err = mongo.NewClient(options.Client().ApplyURI(db.DbServer))
	if db.Err != nil {
		return false
	}
	db.Ctx, db.Cancel = context.WithTimeout(context.Background(), 30*time.Second)
	if db.Err != nil {
		return false
	}
	if 	db.Err = db.Client.Connect(db.Ctx);db.Err != nil {
		return false
	}
	return true
}

func (db *Database) Disconnect() {
	defer db.Cancel()
	defer func() {
		if db.Err = db.Client.Disconnect(db.Ctx); db.Err != nil {
			panic(db.Err)
		}
	}()
}

func (db *Database) Ping() bool {
	if db.Err = db.Client.Ping(db.Ctx, readpref.Primary()); db.Err != nil {
		return false
	}
	return true
}

func (db *Database) QueryOne(database string, col string, filter interface{}) (bson.M) {
	_col := db.Client.Database(database).Collection(col)
	db.Ctx, db.Cancel = context.WithTimeout(context.Background(), 10*time.Second)
	var result bson.M
	db.Err = _col.FindOne(db.Ctx, filter).Decode(&result)
	return result
}

func (db *Database) InsertOne(database string, col string, doc interface{}) (interface{}) {
	_col := db.Client.Database(database).Collection(col)
	var result interface{}
	db.Ctx, db.Cancel = context.WithTimeout(context.Background(), 10*time.Second)
	result, db.Err = _col.InsertOne(db.Ctx, doc)
	return result
}

func (db* Database) InsertMany(database string, col string, doc []interface{})(interface{}){
	_col := db.Client.Database(database).Collection(col)
	db.Ctx, db.Cancel = context.WithTimeout(context.Background(), 10*time.Second)
	var result interface{}
	result, db.Err = _col.InsertMany(db.Ctx, doc)
	return result
}

func (db *Database) QueryMany(database string, col string,filter interface{}) ([]bson.D) {
	_col := db.Client.Database(database).Collection(col)
	db.Ctx, db.Cancel = context.WithTimeout(context.Background(), 10*time.Second)	
	cusor, err := _col.Find(db.Ctx,filter)
	result:=[] bson.D{}
	if err==nil{
		err=cusor.All(db.Ctx,&result)
	}
	db.Err=err
	return result
}
### Database Go Module
This is a simple go module that allows you to interact with the MongoDB using mongo-driver. The module is generic and can be used in any go project.

#### Usage
1. Create an instance of the Database structure.

`var database = Database{DbName: "example", DbServer: "mongodb://localhost:27017"}`

Ensure the path. Ensure that the **DbName** and **DbServer** are initialized. failure to which the **TestRequiredString** will fail.

2. Making queries.

`filter := bson.M{}`
` database.QueryOne(database.DbName, "users", filter)`

or

`filter := bson.M{"$or":[]bson.M{{"name":"test"},{"name":"test1"}}}`
`database.QueryMany(database.DbName, "users", filter)`

The **QueryMany** and **QueryOne** requires that you pass in the database name **DbName** , the collection and a filter which is a **bson object**.

3. Inserting records.

`persons := []interface{}{
				Person{Name: "test", Address: Address{City: "NewYork"}},
				Person{Name: "test1", Address: Address{City: "Nairobi"}},
			}`
` database.InsertMany(database.DbName,"users",persons)`

or

`person := Person{Name: "test", Address: Address{City: "NewYork"}}`
`database.InsertOne(database.DbName, "users", person)`

The **InsertMany** and **InsertOne** requires that you pass in the database name **DbName** , the collection and a filter which is a **bson object**.

4. Delete records.

`filter := bson.M{"$or":[]bson.M{{"name":"test"},{"name":"test1"}}}`
`var options *options.DeleteOptions`
`database.DeleteOne(database.DbName,"users",filter,options)`

or

`filter := bson.M{"$or":[]bson.M{{"name":"test"},{"name":"test1"}}}`
`var options *options.DeleteOptions`
`database.DeleteMany(database.DbName,"users",filter,options)`

The **DeleteOne** and **Deletemany** requires that you pass in the database name **DbName** , the collection and a filter which is a **bson object** and additional options.

5. Update records.

`filter:=bson.M{}`
`update:=bson.M{"$set":bson.M{"name":"test2"}}`
`var options *options.UpdateOptions`
`database.UpdateMany(database.DbName,"users",filter,update,options)`

or

`filter:=bson.M{"name":"test"}`
`update:=bson.M{"$set":bson.M{"name":"test1"}}`
`var options *options.UpdateOptions`
`database.UpdateOne(database.DbName,"users",filter,update,options)`

The **UpdateOne** and **UpdateMany** requires that you pass in the database name **DbName** , the collection and a filter which is a **bson object** and additional options.

6. Distinct records.

`var filter string="name"`
`options:=  options.Distinct().SetMaxTime(2 * time.Second)`
`result:=database.FindDistinct(database.DbName,"users",filter,options)`

The **FindDistinct** requires that you pass in the database name **DbName** , the collection and a filter which is a **bson object** and additional options.	
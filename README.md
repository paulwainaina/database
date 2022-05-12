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

The **QueryMany** and **QueryOne** requires that you pass in the database name **DbName** , the collection and a filter which is a **bson object**

3. Inserting records.

`persons := []interface{}{
				Person{Name: "test", Address: Address{City: "NewYork"}},
				Person{Name: "test1", Address: Address{City: "Nairobi"}},
			}`
` database.InsertMany(database.DbName,"users",persons)`

or

`person := Person{Name: "test", Address: Address{City: "NewYork"}}`
`database.InsertOne(database.DbName, "users", person)`

The **InsertMany** and **InsertOne** requires that you pass in the database name **DbName** , the collection and a filter which is a **bson object**
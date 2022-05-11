package database

import (
	"testing"
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

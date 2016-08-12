package main

import (
	"testing"
	"time"
)

// TestPutNotification ...
func TestPutNotification(t *testing.T) {
	id := 3
	db := DBClient()
	PutNotification(db, id)
	table := db.Table("tool_github_notifications")
	var result Notification
	err := table.Get("id", id).One(&result)
	if err != nil {
		t.Errorf("An error occurred. %s", err)
		return
	}
	if result.ID != 3 {
		t.Errorf("data is wrong. expected: %d, actual: %d", id, result.ID)
	}
}

func TestCheckNotificationExistence(t *testing.T) {
	id := 5
	db := DBClient()
	table := db.Table("tool_github_notifications")
	timeLayout := "Mon Jan 02 15:04:05 GMT+0000"
	table.Put(Notification{ID: id, CreatedAt: time.Now().Format(timeLayout)}).Run()
	exist, err := CheckNotificationExistence(db, id)
	if err != nil {
		t.Errorf("An error occurred. %s", err)
		return
	}
	if exist == false {
		t.Errorf("Record doesn't exist.")
	}
}

func TestCheckNotificationNoExistence(t *testing.T) {
	id := 7
	db := DBClient()
	exist, err := CheckNotificationExistence(db, id)
	if err != nil {
		t.Errorf("An error occurred. %s", err)
		return
	}
	if exist == true {
		t.Errorf("Record exists")
	}
}

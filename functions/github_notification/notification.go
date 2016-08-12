package main

import (
	"log"
	"time"

	"github.com/guregu/dynamo"
)

// Notification - table tool_github_notifications
type Notification struct {
	ID        int    `dynamo:"id,hash"`
	CreatedAt string `dynamo:"created_at"`
}

const tableName = "tool_github_notifications"

// PutNotification - Put notification record
func PutNotification(db *dynamo.DB, id int) error {
	table := db.Table(tableName)
	timeLayout := "Mon Jan 02 15:04:05 GMT+0000"
	if err := table.Put(Notification{ID: id, CreatedAt: time.Now().Format(timeLayout)}).Run(); err != nil {
		log.Println(err)
		return err
	}
	return nil
}

// CheckNotificationExistence - Check record existence
func CheckNotificationExistence(db *dynamo.DB, id int) (bool, error) {
	table := db.Table(tableName)
	cnt, err := table.Get("id", id).Count()
	if err != nil {
		return false, err
	}
	return cnt > 0, nil
}

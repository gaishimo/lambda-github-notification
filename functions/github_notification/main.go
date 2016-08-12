package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/apex/go-apex"
	"github.com/google/go-github/github"
	"github.com/guregu/dynamo"
)

type message struct {
	Target string `json:"target"`
}

func main() {
	apex.HandleFunc(func(event json.RawMessage, ctx *apex.Context) (interface{}, error) {
		var m message
		if err := json.Unmarshal(event, &m); err != nil {
			return nil, err
		}
		Proceed()
		return m, nil
	})
}

// NotificationRecord ...
type NotificationRecord struct {
	Notification github.Notification
	Comment      *github.IssueComment
	PullRequest  *github.PullRequest
}

// Proceed ...
func Proceed() error {
	tokenClient := MakeTokenClient()
	db := DBClient()
	githubNotifications, err := FetchNotifications(tokenClient, true)
	if err != nil {
		log.Println(err)
		return err
	}
	var records = []NotificationRecord{}
	for _, githubNotification := range githubNotifications {
		notificationID, _ := strconv.Atoi(*githubNotification.ID)
		alreadyExists, err := CheckNotificationExistence(db, notificationID)
		if err != nil {
			log.Println(err)
			return err
		}
		if alreadyExists == false {
			record, _ := makeNotificationRecord(db, tokenClient, githubNotification)
			records = append(records, record)
		}
	}

	Notify(records)
	return nil
}

func makeNotificationRecord(db *dynamo.DB, tokenClient *http.Client, notification github.Notification) (NotificationRecord, error) {
	notificationID, _ := strconv.Atoi(*notification.ID)
	var notificationRecord NotificationRecord
	var commentURL string
	var comment *github.IssueComment
	var pullRequest *github.PullRequest

	if notification.Subject.LatestCommentURL != nil {
		commentURL = *notification.Subject.LatestCommentURL
		comment, _ = FetchComment(tokenClient, commentURL)
	}
	if comment == nil {
		pullRequest, _ = FetchPullRequest(tokenClient, *notification.Subject.URL)
	}

	PutNotification(db, notificationID)
	notificationRecord = NotificationRecord{notification, comment, pullRequest}
	return notificationRecord, nil
}

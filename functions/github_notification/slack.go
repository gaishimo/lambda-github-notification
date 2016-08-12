package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
)

// Notify ...
func Notify(records []NotificationRecord) (string, error) {
	webhookURL := os.Getenv("SLACK_WEBHOOK_URL")
	if webhookURL == "" {
		webhookURL = "https://hooks.slack.com/services/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"
	}

	attachments := []attachment{}
	for _, record := range records {
		title := *record.Notification.Repository.Name + " " + *record.Notification.Subject.Title
		if record.Comment != nil {
			attachments = append(attachments, attachment{
				Title:     title,
				TitleLink: *record.Comment.HTMLURL,
				Text:      *record.Comment.Body,
				ThumbURL:  *record.Comment.User.AvatarURL,
				Color:     "#4C4C4C",
			})
		} else if record.PullRequest != nil {
			attachments = append(attachments, attachment{
				Title:     title,
				TitleLink: *record.PullRequest.HTMLURL,
				Text:      "...",
				ThumbURL:  *record.PullRequest.User.AvatarURL,
				Color:     "#4C4C4C",
			})
		}
	}
	params, _ := json.Marshal(data{
		"Github Notification", ":speech_balloon:", attachments,
	})
	resp, err := http.PostForm(
		webhookURL,
		url.Values{"payload": {string(params)}},
	)
	if err != nil {
		return "", err
	}
	body, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	return string(body), nil
}

type data struct {
	Username   string       `json:"username"`
	IconEmoji  string       `json:"icon_emoji"`
	Attachment []attachment `json:"attachments"`
}

type attachment struct {
	Title     string `json:"title"`
	TitleLink string `json:"title_link"`
	Text      string `json:"text"`
	ThumbURL  string `json:"thumb_url"`
	Color     string `json:"color"`
	ImageURL  string `json:"image_url"`
}

package main

import (
	"net/http"
	"os"
	"regexp"
	"strconv"
	"time"

	"golang.org/x/oauth2"

	"github.com/google/go-github/github"
)

const defaultToken = "XXXXXXXXXXXXXXX"

// FetchNotifications ..
func FetchNotifications(c *http.Client, participating bool) ([]github.Notification, error) {
	client := github.NewClient(c)
	sinceTime := time.Now().Add(-(time.Duration(72) * time.Hour))
	opt := &github.NotificationListOptions{
		All:           false,
		Participating: participating,
		Since:         sinceTime,
		Before:        time.Now(),
	}
	notifications, _, err := client.Activity.ListNotifications(opt)
	if err != nil {
		return nil, err
	}
	return notifications, nil
}

// ScanValuesFromCommentURL return array of [url, owner, repo, id]
func ScanValuesFromCommentURL(url string) []string {
	assigned := regexp.MustCompile(`https:\/\/api.github.com\/repos\/(\S+)/(\S+)\/issues\/comments\/(\d+)`)
	results := assigned.FindStringSubmatch(url)
	return results
}

// ScanValuesFromPullRequestURL ...
func ScanValuesFromPullRequestURL(url string) []string {
	assigned := regexp.MustCompile(`https:\/\/api.github.com\/repos\/(\S+)\/(\S+)\/pulls\/(\d+)`)
	results := assigned.FindStringSubmatch(url)
	return results
}

// FetchComment ..
func FetchComment(c *http.Client, commentURL string) (*github.IssueComment, error) {
	client := github.NewClient(c)
	values := ScanValuesFromCommentURL(commentURL)
	if len(values) == 0 {
		return nil, nil
	}
	commentID, _ := strconv.Atoi(values[3])
	comment, _, err := client.Issues.GetComment(values[1], values[2], commentID)
	if err != nil {
		return nil, err
	}
	return comment, nil
}

// FetchPullRequest ...
func FetchPullRequest(c *http.Client, pullRequestURL string) (*github.PullRequest, error) {
	client := github.NewClient(c)
	values := ScanValuesFromPullRequestURL(pullRequestURL)
	pullRequestID, _ := strconv.Atoi(values[3])
	pullRequest, _, err := client.PullRequests.Get(values[1], values[2], pullRequestID)
	if err != nil {
		return nil, err
	}
	return pullRequest, nil
}

// MakeTokenClient ...
func MakeTokenClient() *http.Client {
	token := defaultToken
	if os.Getenv("GITHUB_TOKEN") != "" {
		token = os.Getenv("GITHUB_TOKEN")
	}
	tokenSource := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tokenClient := oauth2.NewClient(oauth2.NoContext, tokenSource)
	return tokenClient
}

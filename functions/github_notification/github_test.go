package main

import "testing"

// TestFetchNotifications ...
func TestFetchNotifications(t *testing.T) {
	tokenClient := MakeTokenClient()
	_, err := FetchNotifications(tokenClient, false)
	if err != nil {
		t.Errorf("An error occurred. %s", err)
	}
}

// TestScanValuesFromCommentURL ...
func TestScanValuesFromCommentURL(t *testing.T) {
	const url = "https://api.github.com/repos/gaishimo-test/garbage1/issues/comments/224496864"
	actual := ScanValuesFromCommentURL(url)
	if len(actual) != 4 {
		t.Errorf("Length is not 4. actual: %d", len(actual))
	}
	if actual[1] != "gaishimo-test" {
		t.Errorf("owner value is wrong. actual: %s", actual[1])
	}
	if actual[2] != "garbage1" {
		t.Errorf("reponame value is wrong. actual: %s", actual[2])
	}
	if actual[3] != "224496864" {
		t.Errorf("id value is wrong. actual: %s", actual[3])
	}
}

// ScanValuesFromPullRequestURL ...
func TestScanValuesFromPullRequestURL(t *testing.T) {
	const url = "https://api.github.com/repos/gaishimo-test/garbage1/pulls/2"
	actual := ScanValuesFromPullRequestURL(url)
	if len(actual) != 4 {
		t.Errorf("Length is not 4. actual: %d", len(actual))
		return
	}
	if actual[1] != "gaishimo-test" {
		t.Errorf("owner value is wrong. actual: %s", actual[1])
		return
	}
	if actual[2] != "garbage1" {
		t.Errorf("reponame value is wrong. actual: %s", actual[2])
		return
	}
	if actual[3] != "2" {
		t.Errorf("id value is wrong. actual: %s", actual[3])
		return
	}
}

// TestFetchComment ...
func TestFetchComment(t *testing.T) {
	const url = "https://api.github.com/repos/gaishimo-test/garbage1/issues/comments/224496864"
	tokenClient := MakeTokenClient()
	comment, err := FetchComment(tokenClient, url)
	if err != nil {
		t.Errorf("An error occurred. %s", err)
	}
	if *comment.ID != 224496864 {
		t.Errorf("ID is wrong. %d", comment.ID)
	}
}

func TestFetchPullRequest(t *testing.T) {
	const url = "https://api.github.com/repos/gaishimo-test/garbage1/pulls/2"
	tokenClient := MakeTokenClient()
	pullRequest, err := FetchPullRequest(tokenClient, url)
	if err != nil {
		t.Errorf("An error occurred. %s", err)
	}
	if *pullRequest.Number != 2 {
		t.Errorf("ID is wrong. %d", pullRequest.Number)
	}
}

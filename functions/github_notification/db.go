package main

import (
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/guregu/dynamo"
)

// ToolGithubNotifications ...
type ToolGithubNotifications struct {
	ID int `dynamo:"id,hash"`
}

// DBClient ...
func DBClient() *dynamo.DB {
	var awsConfig aws.Config
	if os.Getenv("APP_ENV") != "prod" && os.Getenv("APP_ENV") != "dev" {
		if os.Getenv("DYNAMODB_ENDPOINT") == "" {
			awsConfig.Endpoint = aws.String("localhost:8000")
		} else {
			awsConfig.Endpoint = aws.String(os.Getenv("DYNAMODB_ENDPOINT"))
		}
		awsConfig.DisableSSL = aws.Bool(true)
	}
	return dynamo.New(session.New(), &awsConfig)
}

// CreateTables ...
func CreateTables() {
	db := DBClient()
	err := db.CreateTable("tool_github_notifications", ToolGithubNotifications{}).Run()
	if err != nil {
		log.Println(err)
	}
}

// DeleteTables ...
func DeleteTables() {
	db := DBClient()
	table := db.Table("tool_github_notifications")
	if err := table.DeleteTable().Run(); err != nil {
		log.Println(err)
	}
}

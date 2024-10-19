package main

import (
	"log"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
)

var db *dynamodb.Client
var sm *secretsmanager.Client

func init() {
	dbClient, err := getDbClient()
	if err != nil {
		log.Println("init() error running getClient(): ")
		log.Fatal(err)
	}
	db = dbClient

	smClient, err := getSmClient()
	if err != nil {
		log.Println("init() error running getSecretsManagerClient(): ")
		log.Fatal(err)
	}
	sm = smClient
}

func main() {
	lambda.Start(router)
}

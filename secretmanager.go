package main

import (
	"context"
	"errors"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
)

func getSmClient() (*secretsmanager.Client, error) {
	sdkConfig, err := config.LoadDefaultConfig(context.TODO())
	secretsManagerClient := secretsmanager.NewFromConfig(sdkConfig)

	return secretsManagerClient, err
}

func getApiKey() (*string, error) {
	input := &secretsmanager.GetSecretValueInput{
		SecretId: aws.String(ApiKeyName),
	}

	result, err := sm.GetSecretValue(context.TODO(), input)
	if err != nil {
		log.Println("getApiKey() error running secMan.GetSecretValue")
		return nil, err
	}

	if result.SecretString == nil {
		return nil, errors.New(ApiKeyName + " secret is empty")
	}

	secret := *result.SecretString

	return &secret, nil
}

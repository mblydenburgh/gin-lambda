package repository

import (
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/guregu/dynamo"
)

func AddUser(item UserItem) (string, error) {
	client := dynamo.New(session.New(), &aws.Config{Region: aws.String("us-east-1")})
	tableName := os.Getenv("TABLE_NAME")

	table := client.Table(tableName)

	putAction := table.Put(item)

	if err := putAction.Run(); err != nil {
		log.Printf("Error performing put operation on table")
		return "", err
	}

	return item.UserId, nil

}

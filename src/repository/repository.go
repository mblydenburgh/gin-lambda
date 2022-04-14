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

func GetUser(id string) (*UserItem, error) {
	log.Printf("Looking up user id %v", id)
	client := dynamo.New(session.New(), &aws.Config{Region: aws.String("us-east-1")})
	tableName := os.Getenv("TABLE_NAME")
	table := client.Table(tableName)

	var result UserItem
	rangeKey := "User#"
	err := table.Get("UserId", id).Range("ModelTypeAndId", dynamo.BeginsWith, rangeKey).One(&result)
	if err != nil {
		log.Println("Error performing getItem operation")
		log.Println(err)
		return nil, err
	}

	log.Println("found car")
	return &result, nil
}

func GetAllUsers() ([]UserItem, error) {

}

func DeleteUser(id string) (string, error) {

}

package arn

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

//MapOpenID map an arn to an open_id, empty if it's not created
func MapOpenID(arn string, table string, region string) (string, error) {
	input := &dynamodb.GetItemInput{
		TableName: aws.String(table),
		Key: map[string]*dynamodb.AttributeValue{
			"arn": {S: aws.String(arn)},
			"key": {S: aws.String("open_id")},
		},
		AttributesToGet: []*string{aws.String("arn"), aws.String("key"), aws.String("value")},
	}

	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(region)},
	)
	if err != nil {
		return "", err
	}
	svc := dynamodb.New(sess)

	item, err := svc.GetItem(input)
	if err != nil {
		return "", err
	}
	if attr, ok := item.Item["value"]; ok && attr.S != nil {
		return *attr.S, nil
	}
	return "", nil
}

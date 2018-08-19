package arn

import (
	"fmt"
	"testing"
)

func TestMapOpenID(test *testing.T) {
	openID, err := MapOpenID("arn::itea::1::platform::client::1", "iteacloud-platform-api-auth-dev-PlatformPropertyDynamoDbTable-PJC9OSBPAM9P", "ap-southeast-1")
	if err != nil {
		fmt.Println(err)
		test.Fail()
	}
	fmt.Println(openID)
}

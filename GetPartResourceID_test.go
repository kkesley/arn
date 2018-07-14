package arn

import (
	"testing"
)

func TestGetPartResourceID(test *testing.T) {
	resourceARN := "arn::itea::1::property::hotel::3::room::4::window::5"
	if GetPartResourceID(resourceARN, "room") != "4" {
		test.Error("room must be 4")
	}
}

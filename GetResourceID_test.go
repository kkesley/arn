package arn

import "testing"

func TestGetResourceID(test *testing.T) {
	resourceARN := "arn::itea::1::property::hotel::3::room::4::window::5"
	if GetResourceID(resourceARN) != "5" {
		test.Error("window must be 5")
	}
}

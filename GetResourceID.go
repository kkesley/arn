package arn

import (
	"strings"
)

//GetResourceID get the id of an ARN (deppest child)
func GetResourceID(resourceARN string) string {
	sections := strings.Split(resourceARN, "::")
	return sections[len(sections)-1]
}

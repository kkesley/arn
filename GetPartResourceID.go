package arn

import "strings"

//GetPartResourceID get resource ID of a given part (any child)
func GetPartResourceID(resourceARN string, part string) string {
	sections := strings.Split(resourceARN, "::")
	found := false
	for _, section := range sections {
		if section == part {
			found = true
		} else if found {
			return section
		}
	}
	return ""
}

package utils

import "strings"

func FindStringInArray(list []string, value string) bool {
	for _, item := range list {
		if strings.Contains(item, value) {
			return true
		}
	}
	return false
}

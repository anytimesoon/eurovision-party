package utils

import "strings"

func Slugify(name string) string {
	splitName := strings.Split(name, " ")
	for i, name := range splitName {
		splitName[i] = strings.ToLower(name)
	}
	slug := strings.Join(splitName, "-")
	return slug
}

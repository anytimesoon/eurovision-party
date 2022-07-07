package utils

import "strings"

func Slugify(name string) string {
	splitName := strings.Split(name, " ")
	slug := strings.Join(splitName, "-")
	return slug
}

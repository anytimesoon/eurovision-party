package utils

import (
	"log"
	"regexp"
	"strings"
)

func Slugify(name string) string {
	re, err := regexp.Compile(`[[:^alnum:]]`)
	if err != nil {
		log.Fatal(err)
	}

	splitName := strings.Split(name, " ")
	finalName := make([]string, 0)
	for _, word := range splitName {
		word = re.ReplaceAllString(word, "")
		if word != "" {
			finalName = append(finalName, strings.ToLower(word))
		}
	}
	slug := strings.Join(finalName, "-")
	return slug
}

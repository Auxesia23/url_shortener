package utils

import "regexp"


func ValidateUrl(url string) bool {
	pattern := `^[a-zA-Z0-9_-]{1,20}$`

	re, err := regexp.Compile(pattern)
	if err != nil {
		return false
	}

	return re.MatchString(url)
}
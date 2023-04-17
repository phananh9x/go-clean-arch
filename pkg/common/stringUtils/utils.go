package stringUtils

import (
	"fmt"
	"net/url"
	"regexp"
	"unicode"
)

// RemoveSpace remove all white spaces in a string
func RemoveSpace(s string) string {
	rr := make([]rune, 0, len(s))
	for _, r := range s {
		if !unicode.IsSpace(r) {
			rr = append(rr, r)
		}
	}
	return string(rr)
}

func BuildUrl(fullUrl string, queries map[string]interface{}) string {
	return fullUrl + "?" + QueryEncode(queries)
}

func QueryEncode(queries map[string]interface{}) string {
	values := url.Values{}
	for k, v := range queries {
		values.Add(k, fmt.Sprint(v))
	}
	return values.Encode()
}

// StripTags Remove html tags from string
func StripTags(content string) string {
	re := regexp.MustCompile(`<(.|\n)*?>`)
	return re.ReplaceAllString(content, "")
}


package re

import "regexp"

func GetMatchedGroups(pattern string, input string) []string {
	re := regexp.MustCompile(pattern)
	groups := re.FindAllStringSubmatch(input, -1)
	submatches := make([]string, len(groups))
	for i, group := range groups {
		submatches[i] = group[0]
	}
	return submatches
}

func GetNumbers(input string) []string {
	pattern := `\d+`
	return GetMatchedGroups(pattern, input)
}
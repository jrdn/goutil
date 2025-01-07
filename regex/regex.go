package regex

import "regexp"

// Match helps with regexp matches, returning the named capturing groups in a map
func Match(re *regexp.Regexp, input string) []map[string]string {
	var results []map[string]string
	names := re.SubexpNames()

	for _, match := range re.FindAllStringSubmatch(input, -1) {
		data := make(map[string]string)
		for groupIdx, group := range match {
			data[names[groupIdx]] = group
		}
		results = append(results, data)
	}

	return results
}

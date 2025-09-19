package issues

import "regexp"

const estimateFormatRegex = `Estimate: ([0-9]+(\.[0-9]+)?) day(s)?`

// ContainsEstimate checks if the input text contains "Estimate: <positive number> days".
func ContainsEstimate(text string) bool {
	pattern := estimateFormatRegex

	hasEstimate, err := regexp.MatchString(pattern, text)
	if err != nil {
		return false
	}

	return hasEstimate
}

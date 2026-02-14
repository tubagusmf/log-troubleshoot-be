package helper

import "regexp"

func ExtractCCCode(text string) string {
	re := regexp.MustCompile(`#([A-Z]{2}\d+)`)
	match := re.FindStringSubmatch(text)

	if len(match) > 1 {
		return match[1]
	}
	return ""
}

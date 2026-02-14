package helper

import "strings"

type ParsedReport struct {
	Project  string
	Station  string
	Part     string
	DeviceID string
	Issue    string
	CodeName string
}

func ParseWhatsAppReport(message string) ParsedReport {
	lines := strings.Split(message, "\n")

	report := ParsedReport{}

	for _, line := range lines {
		line = strings.TrimSpace(line)

		switch {
		case strings.HasPrefix(strings.ToLower(line), "project"):
			report.Project = getValue(line)

		case strings.HasPrefix(strings.ToLower(line), "stasiun"):
			report.Station = getValue(line)

		case strings.HasPrefix(strings.ToLower(line), "part"):
			report.Part = getValue(line)

		case strings.HasPrefix(strings.ToLower(line), "id"):
			report.DeviceID = getValue(line)

		case strings.HasPrefix(strings.ToLower(line), "permasalahan"):
			report.Issue = getValue(line)

		case strings.HasPrefix(line, "#"):
			report.CodeName = strings.TrimPrefix(line, "#")
		}
	}

	return report
}

func getValue(line string) string {
	parts := strings.SplitN(line, ":", 2)
	if len(parts) == 2 {
		return strings.TrimSpace(parts[1])
	}
	return ""
}

package mlog

import (
	"fmt"
	"strings"
)

func RemoveComments(lines []string) ([]string, error) {
	var result []string
	for _, line := range lines {
		if strings.HasPrefix(line, "#") {
			continue
		}
		result = append(result, line)
	}
	return result, nil
}

func RemoveEmpty(lines []string) ([]string, error) {
	var result []string
	for _, line := range lines {
		trimmedLine := strings.TrimSpace(line)
		if len(trimmedLine) == 0 {
			continue
		}
		result = append(result, trimmedLine)
	}
	return result, nil
}

func ResolveLabels(lines []string) ([]string, error) {
	labels := make(map[string]int)

	// parse labels
	for i, line := range lines {
		if !strings.Contains(line, ":") {
			continue
		}
		parts := strings.Split(line, ":")
		if len(parts) != 2 {
			return nil, fmt.Errorf("invalid line: could not countain more than two ':' symbols: '%s'", line)
		}
		label := parts[0]
		labels[label] = i
		lines[i] = strings.TrimSpace(parts[1])
	}

	// resolve jumps
	for i, line := range lines {
		if !strings.HasPrefix(line, "jump ") {
			continue
		}
		parts := strings.Split(line, " ")
		_, ok := labels[parts[1]]
		if !ok {
			continue
		}
		parts[1] = fmt.Sprintf("%d", labels[parts[1]])
		lines[i] = strings.Join(parts, " ")
	}
	return lines, nil
}

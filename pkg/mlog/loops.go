package mlog

import (
	"strconv"
	"strings"
)

func ResolveLoops(lines []string) ([]string, error) {
	var err error
	var hadChanges bool

	for i := 0; i < 100; i++ {
		lines, hadChanges, err = resolveNextLoop(lines)
		if err != nil {
			return nil, err
		}
		if !hadChanges {
			return lines, nil
		}
	}

	return lines, nil
}

func resolveNextLoop(lines []string) ([]string, bool, error) {
	var loopStart int
	var loopEnd int
	var loopVar string
	var repetitions uint64
	var err error

	// find the last loop
	for i, line := range lines {
		if strings.HasPrefix(line, "STARTLOOP") {
			parts := strings.Split(line, " ")
			loopVar = parts[1]
			repetitions, err = strconv.ParseUint(parts[2], 10, 0)
			if err != nil {
				return nil, false, err
			}
			loopStart = i

		}
	}

	if loopVar == "" {
		return lines, false, nil
	}

	lines[loopStart] = "" //clear loopstart line

	//find end of this loop
	for i, line := range lines {
		if strings.HasPrefix(line, "ENDLOOP "+loopVar) {
			loopEnd = i
			lines[i] = ""
		}
	}

	if loopEnd == 0 {
		return lines, false, nil
	}

	result := make([]string, 0)

	result = append(result, lines[0:loopStart]...)

	loopLines := lines[loopStart+1 : loopEnd]
	for i := 0; i < int(repetitions); i++ {
		var replacedLines = make([]string, len(loopLines))
		for lineNum, line := range loopLines {
			replacedLines[lineNum] = strings.ReplaceAll(line, "$"+loopVar, strconv.Itoa(i))
		}
		result = append(result, replacedLines...)
	}

	result = append(result, lines[loopEnd+1:]...)
	return result, true, nil
}

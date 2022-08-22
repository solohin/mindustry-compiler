package mlog

import "strings"

type Transformer func(lines []string) ([]string, error)

func Transform(lines []string, transformers ...Transformer) ([]string, error) {
	var err error
	for _, transformer := range transformers {
		lines, err = transformer(lines)
		if err != nil {
			return nil, err
		}
	}
	return lines, nil
}

func Compile(text string) (string, error) {
	transformers := make([]Transformer, 0)
	transformers = append(transformers, RemoveEmpty)
	transformers = append(transformers, RemoveComments)
	transformers = append(transformers, RemoveEmpty)
	transformers = append(transformers, ResolveLoops)
	transformers = append(transformers, ResolveLabels)

	lines := strings.Split(strings.ReplaceAll(text, "\r\n", "\n"), "\n")
	resultLines, err := Transform(lines, transformers...)
	if err != nil {
		return "", err
	}
	return strings.Join(resultLines, "\n"), nil
}

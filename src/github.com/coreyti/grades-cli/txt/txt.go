package txt

import "strings"

func Indent(spaces int, v string) string {
	pad := strings.Repeat(" ", spaces)
	return pad + strings.Replace(v, "\n", "\n"+pad, -1)
}

func Wrap(column int, text string) string {
	if column < 1 {
		return text
	}

	text = strings.TrimSpace(text)
	source := strings.Fields(text)
	result := []string{""}

	var word string
	var i int

	for len(source) > 0 {
		word, source = source[0], source[1:]

		if (len(result[i]) + len(word)) < column {
			result[i] = strings.Join([]string{result[i], word}, " ")
		} else {
			i += 1
			result = append(result, word)
		}
	}

	return strings.TrimSpace(strings.Join(result, "\n"))
}

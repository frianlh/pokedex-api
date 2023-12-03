package form

import "regexp"

// SQLInjector is
func SQLInjector(input string) (output string) {
	re := regexp.MustCompile(`['\"\n\r\t\;\$\^\*\\]|://`)
	output = re.ReplaceAllLiteralString(input, "")
	return output
}

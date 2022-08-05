package util

import (
	"regexp"
	"strings"
)

func CompilePatternWithJavaStyle(pattern string) *regexp.Regexp {
	// compile java、c# named capturing groups style
	if strings.Contains(pattern, "?<") {
		pattern = strings.ReplaceAll(pattern, "?<", "?P<")
	}
	return regexp.MustCompile(pattern)
}

func MatchGroup(pattern string, context string) (paramsMap map[string]string) {
	compRegEx := CompilePatternWithJavaStyle(pattern)
	return MatchGroupWithRegex(compRegEx, context)
}

func MatchGroupWithRegex(compRegEx *regexp.Regexp, context string) (paramsMap map[string]string) {
	match := compRegEx.FindStringSubmatch(context)
	l := len(match)
	if l == 0 {
		return
	}
	paramsMap = make(map[string]string, l)
	for i, name := range compRegEx.SubexpNames() {
		if i > 0 && i <= l {
			paramsMap[name] = match[i]
		}
	}
	return
}

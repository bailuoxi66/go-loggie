package util

import "strings"

const sep = "."

func GetQueryPaths(query string) []string {
	paths := strings.Split(query, sep)
	return paths
}

func GetQueryUpperPaths(query string) ([]string, string) {
	paths := strings.Split(query, sep)
	if len(paths) < 2 {
		return []string{}, query
	}
	upper := paths[:len(paths)-1]
	last := paths[len(paths)-1:]
	lastQuery := last[0]

	return upper, lastQuery
}

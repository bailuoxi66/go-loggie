package util

func Contain(key string, array []string) bool {
	if len(array) > 0 {
		for _, s := range array {
			if s == key {
				return true
			}
		}
	}
	return false
}

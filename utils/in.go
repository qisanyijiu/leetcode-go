package utils

func InStrings(val string, arr []string) bool {
	for _, item := range arr {
		if val == item {
			return true
		}
	}
	return false
}

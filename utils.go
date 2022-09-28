package git4humans

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}

func remove(s []string, str string) []string {
	for i, v := range s {
		if v == str {
			return append(s[:i], s[i+1:]...)
		}
	}
	return s
}

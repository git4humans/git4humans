package git4humans

func indexOf(element string, data []string) int {
	for index, v := range data {
		if element == v {
			return index
		}
	}
	return -1
}

func contains(element string, data []string) bool {
	index := indexOf(element, data)
	return index >= 0
}

func remove(data []string, element string) []string {
	index := indexOf(element, data)

	if index >= 0 {
		return append(data[:index], data[index+1:]...)
	}

	return data
}

package checkers

func StringElementExists(element string, array []string) bool {
	for _, value := range array {
		if value == element {
			return true
		}
	}
	return false
}


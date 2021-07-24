package lib

func ContainStr(item string, list []string) bool {
	for _, v := range list {
		if v == item {
			return true
		}
	}

	return false
}

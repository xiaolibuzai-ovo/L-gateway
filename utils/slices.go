package utils

func ContainsString(vv []string, v string) bool {
	for _, item := range vv {
		if item == v {
			return true
		}
	}
	return false
}

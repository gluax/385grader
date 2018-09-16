package utils

func FilterStringSlice(list []string, fil func(string) bool) []string {
	c := list[:0]

	for _, item := range list {
		if fil(item) {
			c = append(c, item)
		}
	}

	return c
}

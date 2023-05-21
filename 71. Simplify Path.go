package golang

func simplifyPath(path string) string {
	result := []string{""}
	for _, part := range strings.Split(path, "/") {
		if part == "." || part == "" {
		} else if part == ".." {
			if len(result) > 1 {
				result = result[:len(result)-1]
			}
		} else {
			result = append(result, part)
		}
	}
	if len(result) == 1 {
		return "/"
	} else {
		return strings.Join(result, "/")
	}
}

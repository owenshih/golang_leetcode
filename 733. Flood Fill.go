package golang

func floodFill(image [][]int, sr int, sc int, color int) [][]int {
	if image[sr][sc] != color {
		fillColor(image, sr, sc, image[sr][sc], color)
	}
	return image
}

func fillColor(image [][]int, m, n, originalColor, color int) {
	if image[m][n] != originalColor {
		return
	}
	image[m][n] = color
	if m+1 < len(image) {
		fillColor(image, m+1, n, originalColor, color)
	}
	if m-1 >= 0 {
		fillColor(image, m-1, n, originalColor, color)
	}
	if n+1 < len(image[0]) {
		fillColor(image, m, n+1, originalColor, color)
	}
	if n-1 >= 0 {
		fillColor(image, m, n-1, originalColor, color)
	}
}

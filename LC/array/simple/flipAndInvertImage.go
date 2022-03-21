package simple

// flipAndInvertImage 832. 翻转图像
func flipAndInvertImage(image [][]int) [][]int {
	n := len(image)
	for i := 0; i < n; i++ {
		l, r := 0, n-1
		for l < r {
			image[i][l] = image[i][l] ^ 1
			image[i][r] = image[i][r] ^ 1
			image[i][l], image[i][r] = image[i][r], image[i][l]

			l++
			r--
		}
		if l == r {
			image[i][l] = image[i][l] ^ 1
		}
	}
	return image
}

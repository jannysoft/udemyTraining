package stringutil

func reverseTwo(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

// reverseTwo is our unexported function
// unexported function starts with small letter
// unexported function can be used by other files inside the package folder
// unexported function is not visible to files outside the folder package
// this demonstrates how unexported function
// can be used by an exported function in the same package
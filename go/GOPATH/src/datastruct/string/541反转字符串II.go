package main

func reverse(str []byte, start int, end int) {
	i := start
	j := end - 1
	for i < j {
		str[i], str[j] = str[j], str[i]
		i++
		j--
	}
}

func reverseStr(s string, k int) string {
	if len(s) == 0 {
		return ""
	}

	str := []byte(s)
	for i := 0; i < len(s); i += 2 * k {
		if i+k <= len(str) {
			reverse(str, i, i+k)
		} else {
			reverse(str, i, len(s))
		}
	}
	return string(str)
}

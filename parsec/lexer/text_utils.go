package lexer

// 是否为字母
func IsAlpha(c rune) bool {
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z')
}

// 是否为数字
func IsDigit(c rune) bool {
	return c >= '0' && c <= '9'
}

// 是否为空白字符
func IsWhite(c rune) bool {
	return c == ' ' || c == '\n' || c == '\r' || c == '\t'
}

// 字符串是否在数组内
func IsInStringArray(arr []string, s string) bool {
	for _, a := range arr {
		if s == a {
			return true
		}
	}
	return false
}

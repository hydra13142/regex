package regex

func hex(l *int, s []byte, n int) int {
	i, j := 0, 0
	for i < n && i < len(s) {
		switch c := s[i]; {
		case c >= 'a' && c <= 'z':
			j = j*16 + int(c-'a'+10)
		case c >= 'A' && c <= 'Z':
			j = j*16 + int(c-'A'+10)
		case c >= '0' && c <= '9':
			j = j*16 + int(c-'0')
		default:
			*l = j
			return i
		}
		i++
	}
	*l = j
	return i
}

func dec(l *int, s []byte, n int) int {
	i, j := 0, 0
	for i < n && i < len(s) {
		c := s[i]
		if c < '0' || c > '9' {
			break
		}
		j = j*10 + int(c-'0')
		i++
	}
	*l = j
	return i
}

func oct(l *int, s []byte, n int) int {
	i, j := 0, 0
	for i < n && i < len(s) {
		c := s[i]
		if c < '0' || c > '7' {
			break
		}
		j = j*8 + int(c-'0')
		i++
	}
	*l = j
	return i
}

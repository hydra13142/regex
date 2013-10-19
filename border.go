package regex

func border(b uint32, s []byte, n int) bool {
	switch b {
	case 'a':
		if n == 0 {
			return true
		}
	case 'A':
		if n != 0 {
			return true
		}
	case 'z':
		if n == len(s) {
			return true
		}
	case 'Z':
		if n != len(s) {
			return true
		}
	case '^':
		if n == 0 {
			return true
		}
		if n > 0 && (s[n-1] == '\r' || s[n-1] == '\n') && s[n] != '\r' && s[n] != '\n' {
			return true
		}
	case '$':
		if n == len(s) {
			return true
		}
		if n > 0 && (s[n] == '\r' || s[n] == '\n') && s[n-1] != '\r' && s[n-1] != '\n' {
			return true
		}
	case 'b':
		if n == 0 || !label.get(s[n-1]) {
			if label.get(s[n]) {
				return true
			}
		} else
			if n == len(s) || !label.get(s[n]) {
				return true
			}
		}
	case 'B':
		if n == 0 || !label.get(s[n-1]) {
			if !label.get(s[n]) {
				return true
			}
		} else
			if n < len(s) && label.get(s[n]) {
				return true
			}
		}
	}
	return false
}

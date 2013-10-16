package regex

func decode(u *unit, s []byte, c *int) int {
	if len(s) == 0 {
		*c = -1
		return 0
	}
	i, n := 1, -1
	switch s[0] {
	case '\\':
		if len(s) == 1 {
			*c = -1
			return 0
		}
		i = 2
		switch s[1] {
		case 0x0:
			*c = -1
			return 0
		case '0':
			n = 0
			u.set(byte(n))
		case 'r':
			n = '\r'
			u.set(byte(n))
		case 'n':
			n = '\n'
			u.set(byte(n))
		case 't':
			n = '\t'
			u.set(byte(n))
		case 'v':
			n = '\v'
			u.set(byte(n))
		case 'f':
			n = '\f'
			u.set(byte(n))
		case 's':
			u.add(space)
		case 'S':
			u.add(not(space))
		case 'c':
			u.add(alpha)
		case 'C':
			u.add(not(alpha))
		case 'd':
			u.add(digit)
		case 'D':
			u.add(not(digit))
		case 'w':
			u.add(label)
		case 'W':
			u.add(not(label))
		case 'x':
			i = 2 + hex(&n, s[2:], 2)
			if i == 2 {
				n = int('x')
			}
			u.set(byte(n))
		default:
			i = 1 + oct(&n, s[1:], 3)
			if i == 1 || n >= 256 {
				i, n = 2, int(s[1])
			}
			u.set(byte(n))
		}
	default:
		n = int(s[0])
		u.set(byte(n))
	}
	*c = n
	return i
}

func encode(c byte) string {
	switch c {
	case 0x0:
		return `\0`
	case '\a':
		return `\a`
	case '\b':
		return `\b`
	case '\r':
		return `\r`
	case '\n':
		return `\n`
	case '\t':
		return `\t`
	case '\v':
		return `\v`
	case '\f':
		return `\f`
	default:
		if c > 31 && c < 127 {
			return string([]byte{c})
		}
		s := []byte{'\\', 'x', 0, 0}
		i := c / 16
		if i > 9 {
			s[2] = i - 10 + 'a'
		} else {
			s[2] = i + '0'
		}
		j := c % 16
		if j > 9 {
			s[2] = j - 10 + 'a'
		} else {
			s[2] = j + '0'
		}
		return string(s)
	}
}

package regex

type token struct {
	typ int
	dtl int
	unt unit
	pr1 int
	pr2 int
}

func certain(t *token, s []byte) int {
	if len(s) < 4 || s[1] != '{' {
		return 0
	}
	if s[2] < '0' || s[2] > '9' {
		return 0
	}
	i := 2 + dec(&t.pr1, s[2:], 4)
	if s[i] != '}' {
		return 0
	}
	return i + 1
}

func limited(t *token, s []byte) int {
	i, a, b := 1, 0, -1
	if s[1] >= '0' && s[1] <= '9' {
		i += dec(&a, s[i:], 4)
	} else if s[1] != ',' {
		return 0
	}
	if s[i] == ',' {
		i++
		if s[i] >= '0' && s[i] <= '9' {
			i += dec(&b, s[i:], 4)
		} else if i == 2 {
			return 0
		}
	} else if i != 1 {
		b = a
	} else {
		return 0
	}
	if s[i] != '}' {
		return 0
	}
	if i+1 < len(s) && s[i+1] == '?' {
		t.dtl = '?'
		i++
	}
	if a > b && b >= 0 {
		t.pr1, t.pr2 = b, a
	} else {
		t.pr1, t.pr2 = a, b
	}
	return i + 1
}

func bracket(u *unit, s []byte) int {
	var a, b bool
	m, n, i, l := -1, 0, 1, 0
	b = false
	if s[1] == '^' {
		a, i = true, 2
	} else {
		a = false
	}
	for ; i < len(s) && s[i] != ']'; i += l {
		switch s[i] {
		case '[':
			x := &unit{}
			l = bracket(x, s[i:])
			if l == 0 {
				return 0
			}
			u.add(x)
			if b {
				u.set('-')
				b = false
			}
			m = -1
		case '-':
			l = 1
			if b {
				if m > '-' {
					for ; m >= '-'; m-- {
						u.set(byte(m))
					}
				} else {
					for ; m <= '-'; m++ {
						u.set(byte(m))
					}
				}
				m, b = -1, false
			} else if m >= 0 {
				b = true
			} else {
				u.set('-')
				m = '-'
			}
		default:
			l = decode(u, s[i:], &n)
			if n < 0 {
				if b {
					u.set('-')
					b = false
				}
				m = -1
			} else {
				if b {
					if m > n {
						for ; m > n; m-- {
							u.set(byte(m))
						}
					} else {
						for ; m < n; m++ {
							u.set(byte(m))
						}
					}
					m, b = -1, false
				} else {
					m = n
				}
			}
		}
	}
	if s[i] != ']' {
		return 0
	}
	if b {
		u.set('-')
	}
	if a {
		u.nega()
	}
	return i + 1
}

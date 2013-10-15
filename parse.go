package regex

func parse(s []byte) []token {
	p := make([]token, 0, 128)
	i, j, k, c := 0, 0, 0, 0
	t := token{}
	for ; len(s) != 0; s = s[j:] {
		t.dtl, t.typ = 0, int(s[0])
		j = 1
		switch s[0] {
		case '(':
			if len(s) > 2 && s[1] == '?' {
				switch s[2] {
				case ':':
					j, t.dtl = 3, ':'
				case '>':
					j, t.dtl = 3, '>'
					k, t.pr1 = k+1, k
				case '=':
					j, t.dtl = 3, '='
					k, t.pr1 = k+1, k
				case '!':
					j, t.dtl = 3, '!'
					k, t.pr1 = k+1, k
				case '<':
					if len(s) > 3 {
						switch s[3] {
						case '=':
							j, t.dtl = 4, '+'
							k, t.pr1 = k+1, k
						case '!':
							j, t.dtl = 4, '-'
							k, t.pr1 = k+1, k
						}
					}
				}
			} else {
				i++
				t.pr1 = i
			}
		case '|', ')':
			/**/
		case '?', '*', '+':
			if len(s) > 1 && s[1] == '?' {
				j, t.dtl = 2, '?'
			}
		case '{':
			j = limited(&t, s)
		case '@', '#':
			j = certain(&t, s)
		case '^', '$':
			t.typ, t.dtl = 'b', t.typ
		case '\\':
			if len(s) > 1 {
				switch s[1] {
				case 'a', 'A':
					fallthrough
				case 'b', 'B':
					fallthrough
				case 'z', 'Z':
					j, t.typ, t.dtl = 2, 'b', int(s[1])
				default:
					t.typ = 'c'
					j = decode(t.unt.zero(), s, &c)
				}
			}
		case '.':
			t.typ = 'c'
			t.unt.full().cls(0)
		case '[':
			t.typ = 'c'
			j = bracket(t.unt.zero(), s)
		default:
			c, t.typ = t.typ, 'c'
			t.unt.zero().set(byte(c))
			break
		}
		if j == 0 {
			j, c, t.typ = 1, t.typ, 'c'
			t.unt.zero().set(byte(c))
		}
		p = append(p, t)
	}
	return p
}

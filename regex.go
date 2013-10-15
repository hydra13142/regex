package regex

type group struct {
	pos int
	len int
}

// A regex complementation of go.
// It can catch / match groups.
// It can match self in expression.
// It can browse forward / backward.
type Regex struct {
	dad plan
	sub []plan
	grp int
}

func (this *Regex) before(g []group, v plan, s []byte, k int) bool {
	j := 0
	c, n, r := v.bgn, k, v.bgn.path
	B := stack{blk: new(block), pnt: 0}
	for {
		p := (*spot)(nil)
		for ; r != nil && p == nil; r = r.next {
			switch r.valu.typ {
			case 'b':
				if border(r.valu.unt[0], s, n) {
					j, p = 0, r.goal
				}
			case 'c':
				if n > 0 && r.valu.unt.get(s[n-1]) {
					j, p = 1, r.goal
				}
			case 'E':
				g[0].pos = n
				g[0].len = k - n
				return true
			}
		}
		if p != nil {
			B.push(c, n, r)
			c, n, r = p, n-j, p.path
		} else {
			for {
				if !B.pop(&c, &n, &r) {
					return false
				}
				if r != nil {
					break
				}
			}
			if r.valu.typ == '(' {
				g[r.valu.unt[0]].len = 0
			}
		}
	}
	return true
}

func (this *Regex) behind(g []group, v plan, s []byte, k int) bool {
	j := 0
	c, n, r := v.bgn, k, v.bgn.path
	B := stack{blk: new(block), pnt: 0}
	for {
		p := (*spot)(nil)
		for ; r != nil && p == nil; r = r.next {
			switch r.valu.typ {
			case 'b':
				if border(r.valu.unt[0], s, n) {
					j, p = 0, r.goal
				}
			case 'c':
				if n < len(s) && r.valu.unt.get(s[n]) {
					j, p = 1, r.goal
				}
			case 'E':
				g[0].pos = k
				g[0].len = n - k
				return true
			}
		}
		if p != nil {
			B.push(c, n, r)
			c, n, r = p, n+j, p.path
		} else {
			for {
				if !B.pop(&c, &n, &r) {
					return false
				}
				if r != nil {
					break
				}
			}
			if r.valu.typ == '(' {
				g[r.valu.unt[0]].len = 0
			}
		}
	}
	return true
}

func (this *Regex) entire(g []group, v plan, s []byte, k int) bool {
	i, j := 0, 0
	u := make([]group, 1)
	c, n, r := v.bgn, k, v.bgn.path
	B := stack{blk: new(block), pnt: 0}
	for {
		p := (*spot)(nil)
		for ; r != nil && p == nil; r = r.next {
			switch r.valu.typ {
			case 'b':
				if border(r.valu.unt[0], s, n) {
					j, p = 0, r.goal
				}
			case 'c':
				if n < len(s) && r.valu.unt.get(s[n]) {
					j, p = 1, r.goal
				}
			case '@':
				i, j = g[r.valu.unt[0]].pos, g[r.valu.unt[0]].len-1
				if n+j >= len(s) {
					break
				}
				for j >= 0 && s[n+j] != s[i+j] {
					j--
				}
				if j == 0 {
					j, p = g[r.valu.unt[0]].len, r.goal
				}
			case '#':
				if this.entire(g, v, s, n) {
					j, p = g[0].len, r.goal
				}
			case '>':
				if this.entire(u, this.sub[r.valu.unt[0]], s, n) {
					j, p = u[0].len, r.goal
				}
			case '=':
				if this.behind(u, this.sub[r.valu.unt[0]], s, n) {
					j, p = 0, r.goal
				}
			case '!':
				if !this.behind(u, this.sub[r.valu.unt[0]], s, n) {
					j, p = 0, r.goal
				}
			case '+':
				if this.before(u, this.sub[r.valu.unt[0]], s, n) {
					j, p = 0, r.goal
				}
			case '-':
				if !this.before(u, this.sub[r.valu.unt[0]], s, n) {
					j, p = 0, r.goal
				}
			case '(':
				g[r.valu.unt[0]].pos = n
				j, p = 0, r.goal
			case ')':
				g[r.valu.unt[0]].len = n - g[r.valu.unt[0]].pos
				j, p = 0, r.goal
			case 'E':
				g[0].pos = k
				g[0].len = n - k
				return true
			}
		}
		if p != nil {
			B.push(c, n, r)
			c, n, r = p, n+j, p.path
		} else {
			for {
				if !B.pop(&c, &n, &r) {
					return false
				}
				if r != nil {
					break
				}
			}
			if r.valu.typ == '(' {
				g[r.valu.unt[0]].len = 0
			}
		}
	}
	return true
}

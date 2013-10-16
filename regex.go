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
	grp int
	dad plan
	sub []plan
}

func (this *Regex) before(g []group, v plan, s []byte, k int) int {
	i, j, p := 0, 0, (*spot)(nil)
	c, n, r := v.bgn, k, v.bgn.path
	B := stack{blk: new(block), pnt: 0}
	for {
		for p = nil; r != nil && p == nil; r = r.next {
			switch r.valu.typ {
			case 'b':
				if border(r.valu.unt[0], s, n) {
					j, p = 0, r.goal
				}
			case 'c':
				if n > 0 && r.valu.unt.get(s[n-1]) {
					j, p = 1, r.goal
				}
			case '@':
				i, j = g[r.valu.unt[0]].pos, g[r.valu.unt[0]].len
				if n-j < 0 {
					break
				}
				k := j - 1
				for t := n - j; k >= 0 && s[t+k] == s[i+k]; {
					k--
				}
				if k < 0 {
					p = r.goal
				}
			case '#':
				if j = this.before(g, v, s, n); j >= 0 {
					p = r.goal
				}
			case 'E':
				return k - n
			}
		}
		if p != nil {
			B.push(c, n, r)
			c, n, r = p, n-j, p.path
			continue
		}
		for {
			if !B.pop(&c, &n, &r) {
				return -1
			}
			if r != nil {
				break
			}
		}
	}
	return 0
}

func (this *Regex) behind(g []group, v plan, s []byte, k int) int {
	i, j, p := 0, 0, (*spot)(nil)
	c, n, r := v.bgn, k, v.bgn.path
	B := stack{blk: new(block), pnt: 0}
	for {
		for p = nil; r != nil && p == nil; r = r.next {
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
				i, j = g[r.valu.unt[0]].pos, g[r.valu.unt[0]].len
				if n+j > len(s) {
					break
				}
				k := j - 1
				for k >= 0 && s[n+k] == s[i+k] {
					k--
				}
				if k < 0 {
					p = r.goal
				}
			case '#':
				if j = this.behind(g, v, s, n); j >= 0 {
					p = r.goal
				}
			case 'E':
				return n - k
			}
		}
		if p != nil {
			B.push(c, n, r)
			c, n, r = p, n+j, p.path
			continue
		}
		for {
			if !B.pop(&c, &n, &r) {
				return -1
			}
			if r != nil {
				break
			}
		}
	}
	return 0
}

func (this *Regex) entire(g []group, v plan, s []byte, k int) bool {
	i, j, p := 0, 0, (*spot)(nil)
	c, n, r := v.bgn, k, v.bgn.path
	B := stack{blk: new(block), pnt: 0}
	for {
		for p = nil; r != nil; r = r.next {
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
				i, j = g[r.valu.unt[0]].pos, g[r.valu.unt[0]].len
				if n+j > len(s) {
					break
				}
				k := j - 1
				for k >= 0 && s[n+k] == s[i+k] {
					k--
				}
				if k < 0 {
					p = r.goal
				}
			case '#':
				if this.entire(g, v, s, n) {
					j, p = g[0].len, r.goal
				}
			case '>':
				if j = this.behind(g, this.sub[r.valu.unt[0]], s, n); j >= 0 {
					p = r.goal
				}
			case '=':
				if this.behind(g, this.sub[r.valu.unt[0]], s, n) >= 0 {
					j, p = 0, r.goal
				}
			case '!':
				if this.behind(g, this.sub[r.valu.unt[0]], s, n) < 0 {
					j, p = 0, r.goal
				}
			case '+':
				if this.before(g, this.sub[r.valu.unt[0]], s, n) >= 0 {
					j, p = 0, r.goal
				}
			case '-':
				if this.before(g, this.sub[r.valu.unt[0]], s, n) < 0 {
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
			if p != nil {
				break
			}
		}
		if p != nil {
			B.push(c, n, r)
			c, n, r = p, n+j, p.path
			continue
		}
		for {
			if !B.pop(&c, &n, &r) {
				return false
			}
			if r.valu.typ == '(' {
				g[r.valu.unt[0]].len = 0
			}
			if r.next != nil {
				break
			}
		}
		r = r.next
	}
	return true
}

package regex

type medi struct {
	ac bool
	pl plan
}

type content struct {
	grp []bool // 捕获的分组
	ptn []medi // 捕获的模式
	sub []plan // 固化及预查
}

func (this *content) atomic(t token) plan {
	s := twin()
	l := head(s.bgn, s.end)
	switch t.typ {
	case 'b':
		l.valu.typ = 'b'
		l.valu.unt[0] = uint32(t.dtl)
	case 'c':
		l.valu.typ = 'c'
		l.valu.unt = t.unt
	case '@':
		if this.grp[t.pr1] {
			l.valu.typ = '@'
			l.valu.unt[0] = uint32(t.pr1)
		} else {
			l.valu.typ = 0
		}
	case '#':
		if t.pr1 == 0 {
			l.valu.typ = '#'
		} else if this.ptn[t.pr1].ac {
			s = same(this.ptn[t.pr1].pl)
		} else {
			l.valu.typ = 0
		}
	}
	return s
}

func (this *content) check(t []token) bool {
	this.grp = nil
	this.ptn = nil
	this.sub = nil
	r := token{'(', 0, unit{}, 0, 0}
	i, v, k, u := 0, 0, 1, 1
	q := r
	for _, p := range t {
		switch p.typ {
		case ')':
			i--
			fallthrough
		case '?', '*', '+', '{', '|':
			if q.typ == '|' || q.typ == '(' {
				return false
			}
		case '(':
			switch p.dtl {
			case 0x0:
				u++
			case ':':
				k++
			default:
				v++
			}
			i++
		}
		if i < 0 {
			return false
		}
		q = p
	}
	if i != 0 {
		return false
	}
	if v != 0 {
		this.sub = make([]plan, v)
	}
	this.ptn = make([]medi, k)
	this.grp = make([]bool, u)
	return true
}

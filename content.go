package regex

type medi struct {
	ac bool
	pl plan
}

type content struct {
	grp []medi // 捕获的分组
	whl int    // 总数
	cnt int    // 当前
	sub []plan // 固化及预查
	ttl int    // 总数
	num int    // 当前
}

func (this *content) atomic(t token) plan {
	if t.typ == '#' && this.grp[t.pr1].ac {
		return same(this.grp[t.pr1].pl)
	}
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
		if this.grp[t.pr1].ac {
			l.valu.typ = '@'
			l.valu.unt[0] = uint32(t.pr1)
		} else {
			l.valu.typ = 0
		}
	case '#':
		if t.pr1 != 0 {
			l.valu.typ = 0
		} else {
			l.valu.typ = '#'
		}
	}
	return s
}

func (this *content) check(t []token) bool {
	this.ttl, this.sub = 0, nil
	this.whl, this.grp = 0, nil
	this.cnt, this.num = 0, 0
	r := token{'(', 0, unit{}, 0, 0}
	q := r
	i, j, v, k := 0, 0, 0, 1
	for _, p := range t {
		switch p.typ {
		case '@':
			if p.pr1 == 0 {
				return false
			}
		case ')':
			i--
			fallthrough
		case '?', '*', '+', '{', '|':
			if q.typ == '|' || q.typ == '(' {
				return false
			}
		case '(':
			switch p.dtl {
			case '>', '=', '!', '+', '-':
				v++
			case 0:
				k++
			}
			i++
			if j < i {
				j = i
			}
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
		this.ttl, this.sub = v, make([]plan, v)
	}
	this.whl, this.grp = k, make([]medi, k)
	return true
}

package regex

func (this *content) option(s plan, t token) plan {
	if t.dtl == '?' {
		head(s.bgn, s.end).valu.typ = 0
	} else {
		tail(s.bgn, s.end).valu.typ = 0
	}
	return s
}

func (this *content) repeat(s plan, t token) plan {
	m := new(spot)
	n := new(spot)
	m.next = s.bgn
	s.end.next = n
	head(s.end, s.bgn).valu.typ = 0
	if t.dtl == '?' {
		head(s.bgn, n).valu.typ = 0
	} else {
		tail(s.bgn, n).valu.typ = 0
	}
	head(m, s.bgn).valu.typ = 0
	s.bgn, s.end = m, n
	return s
}

func (this *content) beyond(s plan, t token) plan {
	m := new(spot)
	n := new(spot)
	m.next = s.bgn
	s.end.next = n
	head(s.end, s.bgn).valu.typ = 0
	if t.dtl == '?' {
		head(s.end, n).valu.typ = 0
	} else {
		tail(s.end, n).valu.typ = 0
	}
	head(m, s.bgn).valu.typ = 0
	s.bgn, s.end = m, n
	return s
}

func (this *content) region(s plan, t token) plan {
	a := plan{}
	b := a
	a.bgn = new(spot)
	a.end = a.bgn
	i, j, k := t.pr1, t.pr2, 0
	if i == j {
		if i != 0 {
			for k = 1; k < i; k++ {
				b = same(s)
				head(a.end, b.bgn).valu.typ = 0
				a.end.next = b.bgn
				a.end = b.end
			}
			head(a.end, s.bgn).valu.typ = 0
			a.end.next = s.bgn
			a.end = s.end
		} else {
			a.end = new(spot)
			a.bgn.next = a.end
			head(a.bgn, a.end).valu.typ = 0
		}
	} else {
		for k = 0; k < i; k++ {
			b = same(s)
			head(a.end, b.bgn).valu.typ = 0
			a.end.next = b.bgn
			a.end = b.end
		}
		if j > 0 {
			n := new(spot)
			for k++; k < j; k++ {
				b = same(s)
				head(a.end, b.bgn).valu.typ = 0
				if t.dtl == '?' {
					head(b.bgn, n).valu.typ = 0
				} else {
					tail(b.bgn, n).valu.typ = 0
				}
				a.end.next = b.bgn
				a.end = b.end
			}
			head(a.end, s.bgn).valu.typ = 0
			if t.dtl == '?' {
				head(s.bgn, n).valu.typ = 0
			} else {
				tail(s.bgn, n).valu.typ = 0
			}
			head(s.end, n).valu.typ = 0
			a.end.next = s.bgn
			s.end.next = n
			a.end = n
		} else if j < 0 {
			n := new(spot)
			b.end.next = n
			head(b.end, b.bgn).valu.typ = 0
			if t.dtl == '?' {
				head(b.end, n).valu.typ = 0
			} else {
				tail(b.end, n).valu.typ = 0
			}
			a.end = n
		}
	}
	return a
}

func (this *content) choice(_ token, a, b plan) plan {
	tail(a.bgn, b.bgn).valu.typ = 0
	head(a.end, b.end).valu.typ = 0
	a.end.next = b.bgn
	a.end = b.end
	return a
}

func (this *content) series(_ token, a, b plan) plan {
	head(a.end, b.bgn).valu.typ = 0
	a.end.next = b.bgn
	a.end = b.end
	return a
}

func (this *content) closed(a token, s plan) (n plan) {
	switch a.dtl {
	case ':':
		n = s
	case '>':
		s.bgn.valu.im++
		head(s.end, nil).valu.typ = 'E'
		reduce(&s, "()>=!+-")
		this.sub[a.pr1] = s
		n = twin()
		l := head(n.bgn, n.end)
		l.valu.typ = byte(a.dtl)
		l.valu.unt[0] = uint32(a.pr1)
	case '=', '!':
		s.bgn.valu.im++
		head(s.end, nil).valu.typ = 'E'
		reduce(&s, "()>=!+-#")
		change(&s)
		this.sub[a.pr1] = s
		n = twin()
		l := head(n.bgn, n.end)
		l.valu.typ = byte(a.dtl)
		l.valu.unt[0] = uint32(a.pr1)
	case '+', '-':
		n = anti(s)
		n.bgn.valu.im++
		head(n.end, nil).valu.typ = 'E'
		reduce(&n, "()>=!+-#")
		change(&n)
		this.sub[a.pr1] = n
		n = twin()
		l := head(n.bgn, n.end)
		l.valu.typ = byte(a.dtl)
		l.valu.unt[0] = uint32(a.pr1)
	default:
		this.grp[a.pr1].pl = same(s)
		this.grp[a.pr1].ac = true
		n.bgn = new(spot)
		n.bgn.next = s.bgn
		l := head(n.bgn, s.bgn)
		l.valu.typ = '('
		l.valu.unt[0] = uint32(a.pr1)
		n.end = new(spot)
		s.end.next = n.end
		l = head(s.end, n.end)
		l.valu.typ = ')'
		l.valu.unt[0] = uint32(a.pr1)
	}
	return n
}

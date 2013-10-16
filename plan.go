package regex

type ios struct {
	im int
	ex int
}

type fac struct {
	typ byte
	unt unit
}

type spot struct {
	next *spot
	path *line
	valu ios
}

type line struct {
	next *line
	goal *spot
	valu fac
}

type plan struct {
	bgn *spot
	end *spot
}

func twin() (p plan) {
	p.bgn = new(spot)
	p.end = new(spot)
	p.bgn.next = p.end
	return p
}

func head(p, q *spot) (l *line) {
	if p != nil {
		l = new(line)
		l.goal = q
		l.next = p.path
		p.path = l
		p.valu.ex++
	}
	if q != nil {
		q.valu.im++
	}
	return
}

func tail(p, q *spot) (l *line) {
	var x *line
	if p != nil {
		x = p.path
		l = new(line)
		l.goal = q
		if x == nil {
			p.path = l
		} else {
			for x.next != nil {
				x = x.next
			}
			x.next = l
		}
		p.valu.ex++
	}
	if q != nil {
		q.valu.im++
	}
	return
}

func same(p plan) plan {
	s := map[*spot]*spot{}
	r := &spot{}
	a := &line{}
	for i := p.bgn; i != nil; i = i.next {
		r.next = new(spot)
		r = r.next
		s[i] = r
		r.valu = i.valu
	}
	for i := p.bgn; i != nil; i = i.next {
		a.next = nil
		b := a
		for j := i.path; j != nil; j = j.next {
			b.next = new(line)
			b = b.next
			b.goal = s[j.goal]
			b.valu = j.valu
		}
		s[i].path = a.next
	}
	return plan{s[p.bgn], s[p.end]}
}

func anti(p plan) plan {
	s := map[*spot]*spot{}
	r := (*spot)(nil)
	for i := p.bgn; i != nil; i = i.next {
		s[i] = new(spot)
		s[i].next = r
		r = s[i]
		r.valu.im, r.valu.ex = i.valu.ex, i.valu.im
	}
	for i := p.bgn; i != nil; i = i.next {
		for j := i.path; j != nil; j = j.next {
			b := new(line)
			b.next = s[j.goal].path
			b.valu = j.valu
			b.goal = s[i]
			s[j.goal].path = b
		}
	}
	return plan{s[p.end], s[p.bgn]}
}

func reduce(self *plan, str string) {
	var (
		p, q, r, s *line
		lp, rp, np *spot
		M          line
		R          spot
		i          int
	)
	for lp = self.bgn; lp != nil; lp = lp.next {
		if lp.valu.im == 0 {
			continue
		}
		M.next = lp.path
		for q, p = &M, M.next; p != nil; q, p = p, p.next {
			if p.valu.typ != 0 {
				for i = 0; i < len(str); i++ {
					if p.valu.typ == str[i] {
						break
					}
				}
			}
			if p.valu.typ == 0 || i < len(str) {
				rp = p.goal
				if lp == rp {
					q.next = p.next
					lp.valu.im--
					lp.valu.ex--
					p = q
				} else if rp.valu.im == 1 {
					for r = rp.path; r.next != nil; r = r.next {
					}
					q.next = rp.path
					r.next = p.next
					rp.path = nil
					lp.valu.ex += rp.valu.ex - 1
					rp.valu.im = 0
					rp.valu.ex = 0
					p = q
				} else if lp.valu.ex == 1 {
					for np = self.bgn; np != nil; np = np.next {
						for r = np.path; r != nil; r = r.next {
							if r.goal == lp {
								r.goal = rp
							}
						}
					}
					q.next = nil
					rp.valu.im += lp.valu.im - 1
					lp.valu.im = 0
					lp.valu.ex = 0
					p = q
				} else {
					for r, s = q, rp.path; s != nil; s = s.next {
						r.next = new(line)
						r = r.next
						*r = *s
						if s.goal != nil {
							s.goal.valu.im++
						}
					}
					r.next = p.next
					lp.valu.ex += rp.valu.ex - 1
					rp.valu.im--
					p = q
				}
				lp.path = M.next
			}
		}
	}
	for lp = self.bgn; lp != nil; lp = lp.next {
		if lp.valu.im == 0 {
			continue
		}
		for p = lp.path; p != nil; p = p.next {
			for r, s = p, p.next; s != nil; r, s = s, s.next {
				if p.valu == s.valu && p.goal == s.goal {
					if s.goal != nil {
						s.goal.valu.im--
					}
					r.next = s.next
					lp.valu.ex--
					s = r
				}
			}
		}
	}
	R.next = self.bgn
	for lp, rp = &R, R.next; rp != nil; lp, rp = rp, rp.next {
		i = rp.valu.im
		if i != 0 {
			for p = rp.path; p != nil; p = p.next {
				if p.goal == rp {
					i--
				}
			}
		}
		if i == 0 {
			lp.next = rp.next
			rp = lp
		}
	}
	self.bgn, self.end = R.next, lp
}

func change(self *plan) {
	for t := self.bgn; t != nil; t = t.next {
		l := &line{next: t.path}
		for q, p := l, l.next; p != nil; q, p = p, p.next {
			if p.valu.typ == 'E' {
				q.next = p.next
				p.next = l.next
				t.path = p
				break
			}
		}
	}
}

func update(self *plan) {
	for t := self.bgn; t != nil; t = t.next {
		for p := t.path; p != nil; p = p.next {
			if p.valu.typ == '(' || p.valu.typ == ')' {
				p.valu.typ = 0
			}
		}
	}
}

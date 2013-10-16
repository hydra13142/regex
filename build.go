package regex

func (this *content) work(t []token) *Regex {
	stk := make([]plan, 1024)
	opr := make([]token, 32)
	i, j, k, p := 0, -1, -1, true
	cls := func(end func(int) bool) {
		for ; k >= 0 && !end(k); k-- {
			j--
			switch opr[k].typ {
			case '&':
				stk[j] = this.series(opr[k], stk[j], stk[j+1])
			case '|':
				stk[j] = this.choice(opr[k], stk[j], stk[j+1])
			}
		}
	}
	for ; i < len(t); i++ {
		if p {
			switch t[i].typ {
			case 'b', 'c', '@', '#':
				j++
				stk[j] = this.atomic(t[i])
				p = false
			case '(':
				k++
				opr[k] = t[i]
			default:
				return nil
			}
		} else {
			switch t[i].typ {
			case '?':
				stk[j] = this.option(stk[j], t[i])
			case '*':
				stk[j] = this.repeat(stk[j], t[i])
			case '+':
				stk[j] = this.beyond(stk[j], t[i])
			case '{':
				stk[j] = this.region(stk[j], t[i])
			case ')':
				cls(func(i int) bool { return opr[i].typ == '(' })
				stk[j] = this.closed(opr[k], stk[j])
				k--
			case '|':
				cls(func(i int) bool { return opr[i].typ == '(' })
				k++
				opr[k] = t[i]
				p = true
			default:
				cls(func(i int) bool { return opr[i].typ != '&' })
				k++
				opr[k] = token{'&', 0, unit{}, 0, 0}
				i, p = i-1, true
			}
		}
	}
	cls(func(i int) bool { return opr[i].typ == '(' })
	ans := new(Regex)
	ans.dad = stk[0]
	ans.dad.bgn.valu.im++
	tail(ans.dad.end, nil).valu.typ = 'E'
	reduce(&ans.dad, "")
	ans.sub = this.sub
	ans.grp = len(this.grp)
	return ans
}

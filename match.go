package regex

// Make a Regex from a string
func Compile(s string) *Regex {
	x := parse([]byte(s))
	if x == nil {
		return nil
	}
	c := content{}
	if !c.check(x) {
		return nil
	}
	r := c.work(x)
	return r
}

// try to find the first match
func (this *Regex) Find(s []byte) []byte {
	if this == nil {
		return nil
	}
	i, g := 0, make([]group, this.grp)
	for ; i < len(s); i++ {
		if this.entire(g, this.dad, s, i) {
			break
		}
	}
	if i >= len(s) {
		return nil
	}
	return s[g[0].pos : g[0].pos+g[0].len]
}

// try to find all match, which don't overlap
func (this *Regex) FindAll(s []byte, n int) (w [][]byte) {
	if this == nil {
		return nil
	}
	a := make([]group, this.grp*2)
	g := a[:this.grp]
	G := a[this.grp:]
	for i, j := 0, 0; i < len(s); {
		copy(g, G)
		if this.entire(g, this.dad, s, i) {
			w = append(w, s[g[0].pos:g[0].pos+g[0].len])
			if j++; j == n {
				break
			}
			i += g[0].len
		} else {
			i++
		}
	}
	return w
}

// try to find the first match with submatch
func (this *Regex) FindSubmatch(s []byte) (w [][]byte) {
	if this == nil {
		return nil
	}
	i, g := 0, make([]group, this.grp)
	for ; i < len(s); i++ {
		if this.entire(g, this.dad, s, i) {
			break
		}
	}
	if i >= len(s) {
		return nil
	}
	w = make([][]byte, this.grp)
	for i, u := range g {
		w[i] = s[u.pos : u.pos+u.len]
	}
	return w
}

// try to find all match with submatch, which don't overlap
func (this *Regex) FindAllSubmatch(s []byte, n int) (w [][][]byte) {
	if this == nil {
		return nil
	}
	a := make([]group, this.grp*2)
	g := a[:this.grp]
	G := a[this.grp:]
	for i, j := 0, 0; i < len(s); {
		copy(g, G)
		if this.entire(g, this.dad, s, i) {
			x := make([][]byte, this.grp)
			for i, u := range g {
				x[i] = s[u.pos : u.pos+u.len]
			}
			w = append(w, x)
			if j++; j == n {
				break
			}
			i += g[0].len
		} else {
			i++
		}
	}
	return w
}

// like Find
func (this *Regex) FindString(s string) string {
	return string(this.Find([]byte(s)))
}

// like FindSubmatch
func (this *Regex) FindStringSubmatch(s string) (w []string) {
	y := this.Find([]byte(s))
	if y == nil {
		return nil
	}
	w = make([]string, this.grp)
	for i, u := range y {
		w[i] = string(u)
	}
	return w
}

// like FindAll
func (this *Regex) FindAllString(s string, n int) (w []string) {
	for _, y := range this.FindAll([]byte(s), n) {
		w = append(w, string(y))
	}
	return w
}

// like FindAllSubmatch
func (this *Regex) FindAllStringSubmatch(s string, n int) (w [][]string) {
	for _, y := range this.FindAllSubmatch([]byte(s), n) {
		x := make([]string, this.grp)
		for j, t := range y {
			x[j] = string(t)
		}
		w = append(w, x)
	}
	return w
}

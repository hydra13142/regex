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

// try to match just at start
func (this *Regex) Test(s []byte) [][]byte {
	if this == nil {
		return nil
	}
	g := make([]group, this.grp)
	if !this.entire(g, this.dad, s, 0) {
		return nil
	}
	x := make([][]byte, this.grp)
	for i, u := range g {
		x[i] = s[u.pos : u.pos+u.len]
	}
	return x
}

// try to find the first match
func (this *Regex) Find(s []byte) [][]byte {
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
	x := make([][]byte, this.grp)
	for i, u := range g {
		x[i] = s[u.pos : u.pos+u.len]
	}
	return x
}

// try to find all match, which don't overlap
func (this *Regex) FindAll(s []byte) [][][]byte {
	if this == nil {
		return nil
	}
	a := make([]group, this.grp*2)
	g := a[:this.grp]
	G := a[this.grp:]
	w := make([][][]byte, 0, 12)
	for i := 0; i < len(s); {
		copy(g, G)
		if this.entire(g, this.dad, s, i) {
			x := make([][]byte, this.grp)
			for i, u := range g {
				x[i] = s[u.pos : u.pos+u.len]
			}
			w = append(w, x)
			i += len(x[0])
		} else {
			i++
		}
	}
	if len(w) == 0 {
		w = nil
	}
	return w
}

// like Test
func (this *Regex) TestString(s string) []string {
	y := this.Test([]byte(s))
	if y == nil {
		return nil
	}
	x := make([]string, this.grp)
	for i, u := range y {
		x[i] = string(u)
	}
	return x
}

// like Find
func (this *Regex) FindString(s string) []string {
	y := this.Find([]byte(s))
	if y == nil {
		return nil
	}
	x := make([]string, this.grp)
	for i, u := range y {
		x[i] = string(u)
	}
	return x
}

// like FindAll
func (this *Regex) FindAllString(s string) [][]string {
	v := this.FindAll([]byte(s))
	if v == nil {
		return nil
	}
	w := make([][]string, len(v))
	for i, y := range v {
		x := make([]string, this.grp)
		for j, t := range y {
			x[j] = string(t)
		}
		w[i] = x
	}
	return w
}

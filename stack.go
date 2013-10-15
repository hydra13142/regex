package regex

type state struct {
	cur *spot
	pos int
	now *line
}

type block struct {
	lst *block
	val [1024]state
}

type stack struct {
	blk *block
	pnt int
}

func (this *stack) push(c *spot, n int, r *line) {
	this.blk.val[this.pnt] = state{c, n, r}
	this.pnt++
	if this.pnt >= 1024 {
		n := new(block)
		n.lst = this.blk
		this.blk = n
		this.pnt = 0
	}
}

func (this *stack) pop(c **spot, n *int, r **line) bool {
	this.pnt--
	if this.pnt < 0 {
		this.blk = this.blk.lst
		this.pnt = 1023
	}
	if this.blk == nil {
		return false
	}
	*c = this.blk.val[this.pnt].cur
	*n = this.blk.val[this.pnt].pos
	*r = this.blk.val[this.pnt].now
	return true
}

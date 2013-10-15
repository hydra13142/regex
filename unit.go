package regex

type unit [8]uint32

var (
	digit unit
	alpha unit
	space unit
	label unit
)

func (u *unit) set(i byte) *unit {
	u[i/32] |= 1 << (uint(i) % 32)
	return u
}

func (u *unit) cls(i byte) *unit {
	u[i/32] &= ^(1 << (uint(i) % 32))
	return u
}

func (u *unit) get(i byte) bool {
	return u[i/32]&(1<<(uint(i)%32)) != 0
}

func (u *unit) add(p *unit) *unit {
	for i := 0; i < 8; i++ {
		(*u)[i] |= (*p)[i]
	}
	return u
}

func (u *unit) addnot(p *unit) *unit {
	for i := 0; i < 8; i++ {
		(*u)[i] |= ^(*p)[i]
	}
	return u
}

func (u *unit) full() *unit {
	for i := 0; i < 8; i++ {
		(*u)[i] = ^uint32(1)
	}
	return u
}

func (u *unit) zero() *unit {
	for i := 0; i < 8; i++ {
		(*u)[i] = 0
	}
	return u
}

func (u *unit) nega() *unit {
	for i := 0; i < 8; i++ {
		(*u)[i] = ^(*u)[i]
	}
	return u
}

func init() {
	for i := 0; i < 8; i++ {
		digit[i] = 0
		alpha[i] = 0
		space[i] = 0
		label[i] = 0
	}
	for i := '0'; i <= '9'; i++ {
		label.set(byte(i))
		digit.set(byte(i))
	}
	for i := 'A'; i <= 'Z'; i++ {
		label.set(byte(i))
		alpha.set(byte(i))
	}
	for i := 'a'; i <= 'z'; i++ {
		label.set(byte(i))
		alpha.set(byte(i))
	}
	label.set('-')
	space.set(' ')
	space.set('\r')
	space.set('\n')
	space.set('\t')
	space.set('\v')
	space.set('\f')
}

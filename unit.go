package regex

type unit [8]uint32

var (
	space = &unit{0x00003e00, 0x00000001, 0x00000000, 0x00000000, 0x00000000, 0x00000000, 0x00000000, 0x00000000}
	digit = &unit{0x00000000, 0x03ff0000, 0x00000000, 0x00000000, 0x00000000, 0x00000000, 0x00000000, 0x00000000}
	alpha = &unit{0x00000000, 0x00000000, 0x07fffffe, 0x07fffffe, 0x00000000, 0x00000000, 0x00000000, 0x00000000}
	label = &unit{0x00000000, 0x03ff0000, 0x87fffffe, 0x07fffffe, 0x00000000, 0x00000000, 0x00000000, 0x00000000}
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
		u[i] |= p[i]
	}
	return u
}

func (u *unit) dec(p *unit) *unit {
	for i := 0; i < 8; i++ {
		u[i] &= ^p[i]
	}
	return u
}

func (u *unit) full() *unit {
	for i := 0; i < 8; i++ {
		u[i] = ^uint32(1)
	}
	return u
}

func (u *unit) zero() *unit {
	for i := 0; i < 8; i++ {
		u[i] = 0
	}
	return u
}

func (u *unit) nega() *unit {
	for i := 0; i < 8; i++ {
		u[i] = ^u[i]
	}
	return u
}

func both(a *unit, b *unit) *unit {
	p := &unit{}
	for i := 0; i < 8; i++ {
		p[i] = a[i] & b[i]
	}
	return p
}

func or(a *unit, b *unit) *unit {
	p := &unit{}
	for i := 0; i < 8; i++ {
		p[i] = a[i] | b[i]
	}
	return p
}

func xor(a *unit, b *unit) *unit {
	p := &unit{}
	for i := 0; i < 8; i++ {
		p[i] = a[i] ^ b[i]
	}
	return p
}

func not(u *unit) *unit {
	p := &unit{}
	for i := 0; i < 8; i++ {
		p[i] = ^u[i]
	}
	return p
}

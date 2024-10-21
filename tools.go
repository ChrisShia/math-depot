package maths

type ModuloCounter struct {
	Start, modulo, i, c int
}

func NewModCounterWithOffset(start, modulo, offset int) *ModuloCounter {
	c := NewModCounter(start, modulo)
	c.Offset(offset)
	return c
}

func NewModCounter(start, mod int) *ModuloCounter {
	return &ModuloCounter{Start: start, modulo: mod}
}

func DefaultModCounter(mod int) *ModuloCounter {
	return &ModuloCounter{modulo: mod}
}

func (mc *ModuloCounter) ModIndex() int {
	return mc.c + mc.Start
}

func (mc *ModuloCounter) Offset(offset int) {
	mc.c = offset % mc.modulo
}

func (mc *ModuloCounter) Next() int {
	mc.i++
	mc.c++
	mc.c = mc.c % mc.modulo
	return mc.ModIndex()
}

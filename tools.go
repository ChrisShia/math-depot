package maths

type ModuloCounter struct {
	start, modulo, i, c int
}

func NewModCounterWithOffset(start, modulo, offset int) *ModuloCounter {
	c := &ModuloCounter{start: start, modulo: modulo}
	c.Offset(offset)
	return c
}

func NewModCounter(start, mod int) *ModuloCounter {
	return &ModuloCounter{start: start, modulo: mod}
}

func DefaultModCounter(mod int) *ModuloCounter {
	return &ModuloCounter{modulo: mod}
}

func (mc *ModuloCounter) ModIndex() int {
	return mc.c + mc.start
}

func (mc *ModuloCounter) Offset(offset int) {
	mc.i += offset
}

func (mc *ModuloCounter) Next() int {
	mc.i++
	mc.c = mc.i % mc.modulo
	return mc.ModIndex()
}

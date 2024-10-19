package maths

type ModuloCounter struct {
	start, modulo, i, c int
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

func (mc *ModuloCounter) Next() int {
	mc.i++
	if mc.i < mc.modulo {
		mc.c = mc.i
		return mc.ModIndex()
	}
	mc.c = mc.i % mc.modulo
	return mc.ModIndex()
}

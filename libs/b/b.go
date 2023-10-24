package b

import "github.com/ramilmsh/golangtest/libs/a"

func (b *B) ToA() *a.A {
	return &a.A{A: b.B}
}

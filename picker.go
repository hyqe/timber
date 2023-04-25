package timber

import (
	"fmt"
	"io"
)

type Picker interface {
	Pick(Lumber)
}

func NewPicker(lvl Level, out io.Writer) Picker {
	return &picker{
		lvl:    lvl,
		Writer: newSyncWriter(out),
	}
}

type picker struct {
	lvl Level
	io.Writer
}

func (p *picker) Pick(l Lumber) {
	if p.lvl <= l.Level() {
		io.Copy(p, l.Chop())
		fmt.Fprint(p, "\n")
	}
}

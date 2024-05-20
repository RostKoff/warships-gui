package gui

import tl "github.com/grupawp/termloop"

type Clickable struct {
	*tl.Entity
	ch       chan<- string
	output   string
	Disabled bool
}

func NewClickable(x, y, w, h int, output string, ch chan<- string) *Clickable {
	return &Clickable{
		Entity: tl.NewEntity(x, y, w, h),
		output: output,
		ch:     ch,
	}
}

func NewClickableOn(obj Physical, output string, ch chan<- string) *Clickable {
	x, y := obj.Position()
	w, h := obj.Size()
	return NewClickable(x, y, w, h, output, ch)
}

func (c *Clickable) Tick(e tl.Event) {
	if c.ch == nil || c.output == "" || c.Disabled {
		return
	}

	switch e.Key {
	case tl.MouseLeft:
		c.processClick(e)
	}
}

func (c *Clickable) processClick(e tl.Event) {
	x, y := c.Position()
	w, h := c.Size()
	if e.MouseX >= x && e.MouseY >= y && e.MouseX < (x+w) && e.MouseY < (y+h) {
		select {
		case c.ch <- c.output:
		default:
			//drop
		}
	}
}

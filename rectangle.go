package gui

import (
	tl "github.com/grupawp/termloop"
)

type rectangle struct {
	*tl.Rectangle
	coord string
	ch    chan<- string
}

func newRectangle(rec *tl.Rectangle) *rectangle {
	return &rectangle{Rectangle: rec}
}

func newClickableRectangle(rec *tl.Rectangle, coord string, ch chan<- string) *rectangle {
	return &rectangle{Rectangle: rec, coord: coord, ch: ch}
}

func (c *rectangle) Tick(e tl.Event) {
	if c.ch == nil || c.coord == "" {
		return
	}

	switch e.Key {
	case tl.MouseLeft:
		c.processClick(e)
	}
}

func (c *rectangle) processClick(e tl.Event) {
	x, y := c.Position()
	w, h := c.Size()
	if e.MouseX >= x && e.MouseY >= y && e.MouseX < (x+w) && e.MouseY < (y+h) {
		select {
		case c.ch <- c.coord:
		default:
			// drop
		}
	}
}

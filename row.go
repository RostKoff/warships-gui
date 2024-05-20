package gui

import tl "github.com/grupawp/termloop"

type row struct {
	columns []*Button
}

func (r *row) Draw(s *tl.Screen) {
	for _, column := range r.columns {
		column.Draw(s)
	}
}
func (r *row) Tick(e tl.Event) {}

func (r *row) Position() (int, int) {
	if len(r.columns) == 0 {
		return 0, 0
	}
	column := r.columns[0]
	return column.Position()
}

func (r *row) Size() (int, int) {
	width := 0
	height := 0
	for _, column := range r.columns {
		w, h := column.Size()
		width += w
		if height < h {
			height = h
		}
	}
	return width, height
}

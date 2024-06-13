//go:build !windows

package gui

import (
	tl "github.com/grupawp/termloop"
)

// Color represents an RGB color.
type Color struct {
	Red   uint8
	Green uint8
	Blue  uint8
}

// NewColor returns a new color. Parameters are red, green and blue values.
func NewColor(r, g, b uint8) Color {
	return Color{Red: r, Green: g, Blue: b}
}

func (c Color) toAttr() tl.Attr {
	return tl.RgbTo256Color(int(c.Red), int(c.Green), int(c.Blue))
}

var (
	White  = Color{Red: 208, Green: 208, Blue: 208}
	Black  = Color{Red: 21, Green: 21, Blue: 21}
	Blue   = Color{Red: 108, Green: 153, Blue: 187}
	Red    = Color{Red: 172, Green: 65, Blue: 66}
	Grey   = Color{Red: 105, Green: 105, Blue: 105}
	Green  = Color{Red: 126, Green: 142, Blue: 0}
	Orange = Color{Red: 255, Green: 191, Blue: 0}
)

package gui

import (
	"github.com/google/uuid"
	tl "github.com/grupawp/termloop"
)

type ButtonConfig struct {
	FgColor     Color
	BgColor     Color
	BorderColor Color
	Width       int
	Height      int
	WithBorder  bool
}

type Button struct {
	id uuid.UUID
	*tl.Rectangle
	cfg *ButtonConfig
	txt *tl.Text
}

func NewButtonConfig() *ButtonConfig {
	return &ButtonConfig{
		FgColor:     White,
		BgColor:     Black,
		BorderColor: White,
		Width:       0,
		Height:      0,
		WithBorder:  false,
	}
}

func NewButton(x, y int, text string, cfg *ButtonConfig) *Button {

	if cfg == nil {
		cfg = NewButtonConfig()
	}

	length := len(text)
	if cfg.Width == 0 {
		cfg.Width = length + 2
	}
	if cfg.Height == 0 {
		cfg.Height = 3
	}

	textX := (cfg.Width - length) / 2
	textY := cfg.Height / 2

	rec := tl.NewRectangle(x, y, cfg.Width, cfg.Height, cfg.BgColor.toAttr())
	txt := tl.NewText(x+textX, y+textY, text, cfg.FgColor.toAttr(), cfg.BgColor.toAttr())

	return &Button{
		id:        uuid.New(),
		Rectangle: rec,
		txt:       txt,
		cfg:       cfg,
	}
}

func (b *Button) Draw(s *tl.Screen) {
	defer b.txt.Draw(s)
	if !b.cfg.WithBorder {
		b.Rectangle.Draw(s)
		return
	}
	w, h := b.Rectangle.Size()
	x, y := b.Rectangle.Position()
	bg := b.cfg.BgColor.toAttr()
	fg := b.cfg.BorderColor.toAttr()
	char := ' '
	for i := 0; i < w; i++ {
		for j := 0; j < h; j++ {
			if i == 0 || i == w-1 {
				char = '█'
			} else if j == 0 {
				char = '▀'
			} else if j == h-1 {
				char = '▄'
			} else {
				char = ' '
			}
			s.RenderCell(x+i, y+j, &tl.Cell{Fg: fg, Bg: bg, Ch: char})
		}
	}
}

func (b *Button) ID() uuid.UUID {
	return b.id
}

func (b *Button) Drawables() []tl.Drawable {
	return []tl.Drawable{b}
}

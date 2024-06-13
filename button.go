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
	id     uuid.UUID
	rec    *tl.Rectangle
	txt    *tl.Text
	border bool
	bColor tl.Attr
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
		id:     uuid.New(),
		rec:    rec,
		txt:    txt,
		border: cfg.WithBorder,
		bColor: cfg.BorderColor.toAttr(),
	}
}

func (b *Button) Tick(e tl.Event) {}

func (b *Button) Draw(s *tl.Screen) {
	defer b.txt.Draw(s)
	if !b.border {
		b.rec.Draw(s)
		return
	}
	w, h := b.rec.Size()
	x, y := b.rec.Position()
	fg := b.bColor
	_, bg := b.txt.Color()
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

func (b *Button) Position() (int, int) {
	return b.rec.Position()
}

func (b *Button) Size() (int, int) {
	return b.rec.Size()
}

func (b *Button) SetBgColor(color Color) {
	attr := color.toAttr()
	fg, _ := b.txt.Color()
	b.rec.SetColor(attr)
	b.txt.SetColor(fg, attr)
}

func (b *Button) SetFgColor(color Color) {
	_, bg := b.txt.Color()
	b.txt.SetColor(color.toAttr(), bg)
}

func (b *Button) SetBorderColor(color Color) {
	b.bColor = color.toAttr()
}

func (b *Button) SetText(text string) {
	b.txt.SetText(text)
	x, y := b.rec.Position()
	w, h := b.rec.Size()
	length := len(text)
	txtX := (w - length) / 2
	txtY := h / 2
	b.txt.SetPosition(x+txtX, y+txtY)
}

func (b *Button) Text() string {
	return b.txt.Text()
}

func (b *Button) Drawables() []tl.Drawable {
	return []tl.Drawable{b}
}

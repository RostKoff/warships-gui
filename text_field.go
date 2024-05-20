package gui

import (
	"sync"

	"github.com/google/uuid"
	tl "github.com/grupawp/termloop"
)

type TextFieldConfig struct {
	BgColor      Color
	FgColor      Color
	ActFgColor   Color
	ActBgColor   Color
	UnfilledChar rune
	MaxLength    int
	InputOn      bool
}

func NewTextFieldConfig() *TextFieldConfig {
	return &TextFieldConfig{
		BgColor:      Black,
		FgColor:      White,
		ActFgColor:   Black,
		ActBgColor:   White,
		UnfilledChar: ' ',
		MaxLength:    0,
		InputOn:      false,
	}
}

type TextField struct {
	*tl.Rectangle
	id     uuid.UUID
	txt    []rune
	cfg    *TextFieldConfig
	active bool
	mx     sync.Mutex
}

func NewTextField(
	x int,
	y int,
	w int,
	h int,
	cfg *TextFieldConfig,
) *TextField {
	if cfg == nil {
		cfg = NewTextFieldConfig()
	}
	if cfg.MaxLength == 0 {
		cfg.MaxLength = w * h
	}
	return &TextField{
		Rectangle: tl.NewRectangle(x, y, w, h, cfg.BgColor.toAttr()),
		id:        uuid.New(),
		txt:       make([]rune, 0),
		cfg:       cfg,
	}
}

// Make Draw function better so it doesn't create new cells every time.
func (in *TextField) Draw(s *tl.Screen) {
	in.mx.Lock()
	defer in.mx.Unlock()
	w, h := in.Size()
	x, y := in.Position()
	maxY := y + h
	var bg, fg tl.Attr
	if in.active {
		bg = in.cfg.ActBgColor.toAttr()
		fg = in.cfg.ActFgColor.toAttr()
	} else {
		bg = in.cfg.BgColor.toAttr()
		fg = in.cfg.FgColor.toAttr()
	}
	currI, currJ := 0, 0
	for _, char := range in.txt {
		if currI >= w {
			currI = 0
			currJ++
		}
		s.RenderCell(x+currI, y+currJ, &tl.Cell{Ch: char, Fg: fg, Bg: bg})
		currI++
	}

	unfilled := tl.Cell{Ch: in.cfg.UnfilledChar, Fg: fg, Bg: bg}
	for j := currJ; y+j < maxY; j++ {
		for ; currI < w; currI++ {
			s.RenderCell(x+currI, y+j, &unfilled)
		}
		currI = 0
	}
}

func (in *TextField) Tick(e tl.Event) {
	if !in.cfg.InputOn {
		return
	}

	x, y := in.Position()
	w, h := in.Size()
	switch e.Key {
	case tl.MouseLeft:
		if e.MouseX >= x && e.MouseY >= y && e.MouseX <= x+w && e.MouseY <= y+h {
			in.active = true
		} else {
			in.active = false
		}
	default:
		if in.active {
			in.handleInput(e.Key, e.Ch)
		}
	}
}

func (in *TextField) handleInput(key tl.Key, char rune) {
	in.mx.Lock()
	defer in.mx.Unlock()
	len := len(in.txt)
	switch key {
	case tl.KeyBackspace | tl.KeyBackspace2:
		if len > 0 {
			in.txt = in.txt[:len-1]
		}
	case tl.KeySpace:
		if len < in.cfg.MaxLength {
			in.txt = append(in.txt, ' ')
		}
	default:
		if char != 0 && len < in.cfg.MaxLength {
			in.txt = append(in.txt, char)
		}
	}
}

func (in *TextField) SetText(text string) {
	in.mx.Lock()
	defer in.mx.Unlock()
	in.txt = []rune(text)
}

func (in *TextField) GetText() string {
	return string(in.txt)
}

func (in *TextField) Drawables() []tl.Drawable {
	return []tl.Drawable{in}
}

func (in *TextField) ID() uuid.UUID {
	return in.id
}

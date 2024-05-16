package gui

import (
	"sync"

	"github.com/google/uuid"
	tl "github.com/grupawp/termloop"
)

type InputConfig struct {
	BgColor      Color
	FgColor      Color
	ActFgColor   Color
	ActBgColor   Color
	UnfilledChar rune
	MaxLength    int
}

func NewInputConfig() *InputConfig {
	return &InputConfig{
		BgColor:      Black,
		FgColor:      White,
		ActFgColor:   Black,
		ActBgColor:   White,
		UnfilledChar: '_',
		MaxLength:    0,
	}
}

type Input struct {
	*tl.Rectangle
	id       uuid.UUID
	usrInput []rune
	cfg      *InputConfig
	active   bool
	mx       sync.Mutex
}

func NewInput(
	x int,
	y int,
	w int,
	h int,
	cfg *InputConfig,
) *Input {
	if cfg == nil {
		cfg = NewInputConfig()
	}
	if cfg.MaxLength == 0 {
		cfg.MaxLength = w * h
	}
	return &Input{
		Rectangle: tl.NewRectangle(x, y, w, h, cfg.BgColor.toAttr()),
		id:        uuid.New(),
		usrInput:  make([]rune, 0),
		cfg:       cfg,
	}
}

func (in *Input) Draw(s *tl.Screen) {
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
	for _, char := range in.usrInput {
		if currI >= w {
			currI = 0
			currJ++
		}
		s.RenderCell(x+currI, y+currJ, &tl.Cell{Ch: char, Fg: fg, Bg: bg})
		currI++
	}

	for j := currJ; y+j < maxY; j++ {
		for ; currI < w; currI++ {
			s.RenderCell(x+currI, y+j, &tl.Cell{Ch: in.cfg.UnfilledChar, Fg: fg, Bg: bg})
		}
		currI = 0
	}
}

func (in *Input) Tick(e tl.Event) {
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

func (in *Input) handleInput(key tl.Key, char rune) {
	in.mx.Lock()
	defer in.mx.Unlock()
	len := len(in.usrInput)
	switch key {
	case tl.KeyBackspace | tl.KeyBackspace2:
		if len > 0 {
			in.usrInput = in.usrInput[:len-1]
		}
	case tl.KeySpace:
		if len < in.cfg.MaxLength {
			in.usrInput = append(in.usrInput, ' ')
		}
	default:
		if char != 0 && len < in.cfg.MaxLength {
			in.usrInput = append(in.usrInput, char)
		}
	}
}

func (in *Input) GetInput() string {
	return string(in.usrInput)
}

func (in *Input) Drawables() []tl.Drawable {
	return []tl.Drawable{in}
}

func (in *Input) ID() uuid.UUID {
	return in.id
}

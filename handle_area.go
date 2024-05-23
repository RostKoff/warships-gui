package gui

import (
	"context"

	"github.com/google/uuid"
	tl "github.com/grupawp/termloop"
)

type HandleArea struct {
	id         uuid.UUID
	clickables []*Clickable
	ch         chan string
}

func NewHandleArea(objs map[string]Physical) *HandleArea {
	clickables := make([]*Clickable, 0)
	ch := make(chan string)
	for key, obj := range objs {
		clickables = append(clickables, NewClickableOn(obj, key, ch))
	}
	return &HandleArea{
		id:         uuid.New(),
		clickables: clickables,
		ch:         ch,
	}
}

func (area *HandleArea) SetClickables(clickables []*Clickable) {
	area.clickables = clickables
}

func (area *HandleArea) ID() uuid.UUID {
	return area.id
}

func (area *HandleArea) Drawables() []tl.Drawable {
	d := []tl.Drawable{}
	for _, c := range area.clickables {
		d = append(d, tl.Drawable(c))
	}
	return d
}

func (area *HandleArea) Listen(ctx context.Context) string {
	select {
	case s := <-area.ch:
		return s
	case <-ctx.Done():
		return ""
	}
}

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
	ch := make(chan string)
	clickables := createClickables(objs, ch)
	return &HandleArea{
		id:         uuid.New(),
		clickables: clickables,
		ch:         ch,
	}
}

func (area *HandleArea) SetClickablesOn(objs map[string]Physical) {
	clickables := createClickables(objs, area.ch)
	area.clickables = clickables
}

func createClickables(objs map[string]Physical, ch chan<- string) (clickables []*Clickable) {
	for key, obj := range objs {
		clickables = append(clickables, NewClickableOn(obj, key, ch))
	}
	return
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

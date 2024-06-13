package gui

import (
	"context"

	"github.com/google/uuid"
	tl "github.com/grupawp/termloop"
)

type HandleArea struct {
	id         uuid.UUID
	clickables map[string]*Clickable
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

func (area *HandleArea) Test() {}

func (area *HandleArea) SetClickablesOn(objs map[string]Physical) {
	clickables := createClickables(objs, area.ch)
	area.clickables = clickables
}

func (area *HandleArea) GetClickable(key string) *Clickable {
	clickable, ok := area.clickables[key]
	if !ok {
		return nil
	}
	return clickable
}

func (area *HandleArea) GetClickables() map[string]*Clickable {
	return area.clickables
}

func createClickables(objs map[string]Physical, ch chan<- string) map[string]*Clickable {
	clickables := make(map[string]*Clickable)
	for key, obj := range objs {
		clickables[key] = NewClickableOn(obj, key, ch)
	}
	return clickables
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

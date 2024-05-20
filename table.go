package gui

import (
	"github.com/google/uuid"
	tl "github.com/grupawp/termloop"
)

type TableConfig struct {
	maxChars int
}

type Table struct {
	id    uuid.UUID
	x, y  int
	cfg   *TableConfig
	tBody [][]string
	rows  []*row
}

func (t *Table) ID() uuid.UUID {
	return t.id
}

func (t *Table) Drawables() []tl.Drawable {
	drawables := make([]tl.Drawable, 0)
	for _, row := range t.rows {
		drawables = append(drawables, row)
	}
	return drawables
}

func NewTableConfig() *TableConfig {
	return &TableConfig{
		maxChars: 15,
	}
}

func NewTable(x, y int, cfg *TableConfig, tBody [][]string) *Table {
	if cfg == nil {
		cfg = NewTableConfig()
	}
	if tBody == nil {
		tBody = make([][]string, 0)
	}

	columnLength := defineColumnLength(tBody)
	rows := make([]*row, 0)
	cellCfg := NewButtonConfig()
	cellCfg.FgColor = White
	for j, rowa := range tBody {
		columns := make([]*Button, 0)
		prevI := 0
		for i, cell := range rowa {
			cellCfg.Width = columnLength[i]
			button := NewButton(x+prevI, y+j*3, cell, cellCfg)
			columns = append(columns, button)
			prevI = prevI + columnLength[i]
		}
		rows = append(rows, &row{columns: columns})
	}

	return &Table{
		x:     x,
		y:     y,
		cfg:   cfg,
		tBody: tBody,
		rows:  rows,
	}
}

func (t *Table) Draw(s *tl.Screen) {
	for _, row := range t.rows {
		row.Draw(s)
	}
}

func defineColumnLength(tBody [][]string) (values []int) {
	values = make([]int, 0)

	if len(tBody) == 0 {
		return
	}
	for i := 0; i < len(tBody[0]); i++ {
		maxChars := 0
		for _, row := range tBody {
			if chars := len(row[i]) + 2; maxChars < chars {
				maxChars = chars
			}
		}
		values = append(values, maxChars)
	}
	return
}

func (t *Table) GetRows() []*row {
	return t.rows
}

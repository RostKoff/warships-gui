package gui

type Physical interface {
	Size() (int, int)
	Position() (int, int)
}

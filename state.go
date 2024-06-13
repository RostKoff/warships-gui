package gui

type State string

const (
	Empty    State = "Empty"
	Hit      State = "Hit"
	Miss     State = "Miss"
	Ship     State = "Ship"
	Emphasis State = "Emphasis"
	Blocked  State = "Blocked"
)

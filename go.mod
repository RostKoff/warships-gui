module github.com/RostKoff/warships-gui/v2

go 1.19

retract (
	v2.1.2
	v2.1.1
	// those versions contain incorrect import
	v2.1.0
)

require (
	github.com/google/uuid v1.3.0
	github.com/grupawp/termloop v0.0.0-20230531144437-277a1cbf4c14

)

require (
	github.com/mattn/go-runewidth v0.0.14 // indirect
	github.com/nsf/termbox-go v1.1.1
	github.com/rivo/uniseg v0.4.4 // indirect
)

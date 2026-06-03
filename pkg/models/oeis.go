package models

type OEISSequence struct {
	ID     string   // A000001
	AltIDs []string // M0098, N0035, etc.

	Sequence []int64 // combined %S %T %U
	Name     string  // %N

	Comments   []string // %C
	References []string // %D
	Links      []string // %H
	Formulas   []string // %F
	CrossRefs  []string // %Y
	Programs   []ProgramEntry
	Examples   []string // %e
	Keywords   []string // %K

	OffsetA int // %O first value
	OffsetB int // %O second value

	Author     string   // %A
	Extensions []string // %E
}

type ProgramEntry struct {
	Language string // Maple, Mathematica, PARI, GAP, Python, etc.
	Code     string
}

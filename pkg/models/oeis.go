package models

// OEISSequence is the model for a sequence in the OEIS
// Read more about the format here: https://oeis.org/eishelp1.html
type OEISSequence struct {
	ID     string   // Identification line, example: A000001
	AltIDs []string // M0098, N0035, etc.

	Sequence []int64 // combined %S %T %U
	Name     string  // Name of the Sequence, example: "Numbers n such that 2^n-1 is prime" %N

	Comments   []string // %C
	References []string // Detailed references%D
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

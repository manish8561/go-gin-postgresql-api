package painkiller
// use for stringer package with go generate command
//go:generate stringer -type=Pill

type Pill int

const (
	Placebo Pill = iota
	Aspirin
	Ibuprofen
	Paracetamol
	Acetaminophen = Paracetamol
)
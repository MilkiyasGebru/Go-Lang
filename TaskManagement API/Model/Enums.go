package Model

type WorkStatus int

const (
	NotStarted WorkStatus = iota
	Pending
	Finished
	Discontinued
)

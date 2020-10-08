package igloo

// Dirtier allows structs to track when they are changed and only apply
// complex changes when there is something new to update.
type Dirtier interface {
	IsDirty() bool
	Clean()
}

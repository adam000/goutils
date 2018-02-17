package deferrer

// Deferrer allows for collecting many defer statements for deferring later on.
type Deferrer struct {
	funcs []func()
}

// Cleanup calls the statements that were added to the Deferrer's stack, in reverse order (the defer
// statement works like a stack: https://tour.golang.org/flowcontrol/13)
func (d *Deferrer) Cleanup() {
	for i := len(d.funcs) - 1; i >= 0; i-- {
		d.funcs[i]()
	}

	d.funcs = nil
}

// Defer adds a function to the Deferrer's stack.
func (d *Deferrer) Defer(f func()) {
	d.funcs = append(d.funcs, f)
}

// New creates a new Deferrer for defering actions on *your* timetable.
func New() Deferrer {
	return Deferrer{}
}

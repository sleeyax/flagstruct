package flag

// Flag is a representation of a command-line flag that you can pass through.
type Flag struct {
	// Name of the flag.
	// It will be automatically determined when omitted.
	Name string

	// Usage description.
	Usage string

	// Default value.
	Value string
}

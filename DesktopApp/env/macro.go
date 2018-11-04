package env

type MacroType int

const (
	// this will simply type some text
	MACRO_TEXT MacroType = iota
	// this will attempt to open an application
	MACRO_OPEN
	// this will execute a sequence of keystrokes
	MACRO_SEQUENCE
	// this will run some code
	// TODO: i have no idea what exactly i want from this and therefore how im gonna implement this yet it needs some
	// TODO: thinking about
	MACRO_CODE
)
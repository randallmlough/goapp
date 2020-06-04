package colors

import (
	"github.com/logrusorgru/aurora"
)

func Yellow(arg interface{}) aurora.Value {
	return aurora.BrightYellow(arg)
}

func Magenta(arg interface{}) aurora.Value {
	return aurora.BrightMagenta(arg)
}

func Green(arg interface{}) aurora.Value {
	return aurora.BrightGreen(arg)
}

func Blue(arg interface{}) aurora.Value {
	return aurora.BrightBlue(arg)
}

func Cyan(arg interface{}) aurora.Value {
	return aurora.BrightCyan(arg)
}

func Red(arg interface{}) aurora.Value {
	return aurora.BrightRed(arg)
}

func Black(arg interface{}) aurora.Value {
	return aurora.Black(arg)
}

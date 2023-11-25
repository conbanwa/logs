package logs

func Red(args ...interface{}) {
	Log.output(L_INFO, i, Dye(0, "red", args...))
}

func Green(args ...interface{}) {
	Log.output(L_INFO, i, Dye(0, "green", args...))
}

func Yellow(args ...interface{}) {
	Log.output(L_INFO, i, Dye(0, "yellow", args...))
}

func Blue(args ...interface{}) {
	Log.output(L_INFO, i, Dye(0, "blue", args...))
}

func Magenta(args ...interface{}) {
	Log.output(L_INFO, i, Dye(0, "magenta", args...))
}

func Cyan(args ...interface{}) {
	Log.output(L_INFO, i, Dye(0, "cyan", args...))
}

func Highlight(color string, args ...interface{}) {
	Log.output(L_INFO, i, Dye(1, color, args...))
}

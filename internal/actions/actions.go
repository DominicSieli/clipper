package actions

func Up(key byte) bool {
	return key == 1 || key == 105
}

func Down(key byte) bool {
	return key == 2 || key == 107
}

func Enter(key byte) bool {
	return key == 3 || key == 13 || key == 108
}

func Escape(key byte) bool {
	return key == 4 || key == 27 || key == 106
}

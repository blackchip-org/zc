package app

import "golang.org/x/sys/windows"

func consoleInit() {
	handle, err := windows.GetStdHandle(uint32(4294967285))
	if err != nil {
		log.Printf("unable to get handle: %v", err)
		return
	}
	if err := windows.SetConsoleMode(handle, 7); err != nil {
		log.Printf("unable to set console mode: %v", err)
	}
}

package main

import "github.com/01-edu/z01"

func PrintStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
}

func CloseDoor(ptrDoor *Door) bool {
	PrintStr("Door Closing...\n")
	ptrDoor.state = "CLOSE"
	return true
}

func IsDoorOpen(Door *Door) bool {
	PrintStr("is the Door opened ?\n")
	return Door.state == "OPEN"
}

func IsDoorClose(ptrDoor *Door) bool {
	PrintStr("is the Door closed ?\n")
	return ptrDoor.state == "CLOSE"
}

func OpenDoor(Door *Door) {
	PrintStr("Door Opening...\n")
	Door.state = "OPEN"
}

type Door struct {
	state string
}

func main() {
	door := &Door{}

	OpenDoor(door)
	if IsDoorClose(door) {
		OpenDoor(door)
	}
	if IsDoorOpen(door) {
		CloseDoor(door)
	}
	if door.state == "OPEN" {
		CloseDoor(door)
	}
}

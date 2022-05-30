package asciianimator

import "fmt"

const (
	esc           = "\x1B"
	moveUpTmpl    = esc + "[%dA"
	moveDownTmpl  = esc + "[%dB"
	moveRightTmpl = esc + "[%dC"
	moveLeftTmpl  = esc + "[%dD"
)

func moveUp(n int) {
	fmt.Printf(moveUpTmpl, n)
}

func moveDown(n int) {
	fmt.Printf(moveDownTmpl, n)
}

func moveRight(n int) {
	fmt.Printf(moveRightTmpl, n)
}

func moveLeft(n int) {
	fmt.Printf(moveLeftTmpl, n)
}

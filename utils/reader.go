package utils

import (
	"bufio"
	"fmt"
	"golang.org/x/sys/unix"
	"os"
	"strings"
)

var origTermios *unix.Termios

func init() {
	var err error
	origTermios, err = unix.IoctlGetTermios(int(os.Stdin.Fd()), unix.TCGETS)
	if err != nil {
		panic(err)
	}
}

func ReadInput() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("> ")
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

func ReadChar() (byte, error) {
	if err := setRawMode(); err != nil {
		return 0, err
	}
	defer restoreMode()

	var b []byte = make([]byte, 1)
	_, err := os.Stdin.Read(b)
	if err != nil {
		return 0, err
	}
	return b[0], nil
}

func setRawMode() error {
	rawTermios := *origTermios
	rawTermios.Lflag &^= unix.ECHO | unix.ICANON
	rawTermios.Cc[unix.VMIN] = 1
	rawTermios.Cc[unix.VTIME] = 0
	return unix.IoctlSetTermios(int(os.Stdin.Fd()), unix.TCSETS, &rawTermios)
}

func restoreMode() error {
	return unix.IoctlSetTermios(int(os.Stdin.Fd()), unix.TCSETS, origTermios)
}

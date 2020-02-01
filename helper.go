package helper

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"

	"github.com/mgutz/ansi"
)

// initial setup
var clear map[string]func()

func init() {
	clear = make(map[string]func())
	clear["linux"] = func() {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	clear["windows"] = func() {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

// helper functions
func PrintColoredDashes(dashCount int, color string) string {
	result := ansi.Color(strings.Repeat("-", dashCount), color)
	return fmt.Sprint(result)
}

// PrintColoredText prints colored text
func PrintColoredText(text string, color string) string {
	result := ansi.Color(text, color)
	return fmt.Sprint(result)
}

// ReadCommand reads commands from cli
func ReadCommand() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	option := scanner.Text()
	return option
}

func ClearScreen() {
	value, ok := clear[runtime.GOOS]
	if ok {
		value()
	} else {
		panic("Your plataform is unsuported! I can't clear terminal screen")
	}
}

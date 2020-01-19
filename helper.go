package helper

import (
	"fmt"
	"github.com/mgutz/ansi"
	"os"
	"os/exec"
	"runtime"
	"strings"
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
func PrintColoredDashes(dashCount int, color string) {
	result := ansi.Color(strings.Repeat("-", dashCount), color)
	fmt.Println(result)
}

func PrintColoredText(text string, color string) {
	result := ansi.Color(text, color)
	fmt.Println(result)
}

func ClearScreen() {
	value, ok := clear[runtime.GOOS]
	if ok {
		value()
	} else {
		panic("Your plataform is unsuported! I can't clear terminal screen")
	}
}

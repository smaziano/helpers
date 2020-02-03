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

// PrintColoredDashes - 'mandatory' on vscode
func PrintColoredDashes(dashCount int, color string) string {
	result := ansi.Color(strings.Repeat("-", dashCount), color)
	return fmt.Sprint(result)
}

// PrintColoredText - 'mandatory' on vscode
func PrintColoredText(text string, color string) string {
	result := ansi.Color(text, color)
	return fmt.Sprint(result)
}

// ReadCommand - 'mandatory' on vscode
func ReadCommand() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	option := scanner.Text()
	return option
}

// ClearScreen - 'mandatory' on vscode
func ClearScreen() {
	value, ok := clear[runtime.GOOS]
	if ok {
		value()
	} else {
		panic("Your plataform is unsuported! I can't clear terminal screen")
	}
}

// ClearScreenOrPreviousScreenText - 'mandatory' on vscode
func ClearScreenOrPreviousScreenText() {
	fmt.Println(PrintColoredDashes(40, "white"))
	fmt.Print(PrintColoredText("(c) ", "blue"))
	fmt.Print(PrintColoredText("Limpar a tela", "white"))
	fmt.Print(PrintColoredText(" | ", "white"))
	fmt.Print(PrintColoredText("(q) ", "blue"))
	fmt.Println(PrintColoredText("Menu anterior", "white"))
	fmt.Print(">> ")
}

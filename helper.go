package helper

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"reflect"
	"runtime"
	"strings"

	"github.com/alexeyco/simpletable"
	"github.com/mgutz/ansi"
)

// initial setup
var clear map[string]func()

func init() {
	clear = make(map[string]func())
	clear["linux"] = func() {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		err := cmd.Run()
		CheckError(err)
	}
	clear["windows"] = func() {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		err := cmd.Run()
		CheckError(err)
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
	fmt.Println(PrintColoredDashes(108, "white"))
	fmt.Print(PrintColoredText("(c) ", "blue"))
	fmt.Print(PrintColoredText("Clear screen", "white"))
	fmt.Print(PrintColoredText(" | ", "white"))
	fmt.Print(PrintColoredText("(q) ", "blue"))
	fmt.Print(PrintColoredText("Quit", "white"))
	fmt.Print(PrintColoredText(" | ", "white"))
	fmt.Print(PrintColoredText("(p) ", "blue"))
	fmt.Println(PrintColoredText("Previous menu", "white"))
	fmt.Print(">> ")
}

// ListMenuItem - 'mandatory' on vscode
func ListMenuItem(command string, description string) {
	fmt.Print(PrintColoredText("("+command+") ", "blue"))
	fmt.Println(PrintColoredText(description, "white"))
}

// Loading - 'mandatory' on vscode
func Loading() {
	fmt.Println(PrintColoredDashes(108, "cyan"))
	fmt.Println(PrintColoredText("Loading...", "cyan"))
	fmt.Println(PrintColoredDashes(108, "cyan"))
}

func WindowHeaderWarning(message string) {
	fmt.Println(PrintColoredDashes(108, "green"))
	fmt.Println(PrintColoredText(message, "yellow"))
	fmt.Println(PrintColoredDashes(108, "green"))
}

func WindowHeaderSuccess(message string) {
	fmt.Println(PrintColoredDashes(108, "green"))
	fmt.Println(PrintColoredText(message, "green"))
	fmt.Println(PrintColoredDashes(108, "green"))
}

func WindowHeaderError(message string) {
	fmt.Println(PrintColoredDashes(108, "green"))
	fmt.Println(PrintColoredText(message, "red"))
	fmt.Println(PrintColoredDashes(108, "green"))
}

// GetTypeAttributes - Get Fields
func GetTypeAttributes(azureResource interface{}) (reflect.Type, int) {
	convertedType := reflect.ValueOf(azureResource).Elem()
	return convertedType.Type(), convertedType.NumField()
}

// GetAllFieldsAndValues - Get All Fields And Values
func GetAllFieldsAndValues(azureResource interface{}) ([]string, []string) {
	var fields []string
	var values []string

	resType, size := GetTypeAttributes(azureResource)
	s := reflect.ValueOf(azureResource).Elem()
	for i := 0; i < size; i++ {
		field := resType.Field(i).Name
		value := fmt.Sprintf("%v", s.Field(i).Interface())
		fields = append(fields, field)
		values = append(values, value)
	}
	return fields, values
}

// GenerateCellsForHeader -
func GenerateCellsForHeader(fields []string, color string) []*simpletable.Cell {
	var cells []*simpletable.Cell
	for i := 0; i < len(fields); i++ {
		field := fields[i]
		cell := simpletable.Cell{Align: simpletable.AlignLeft, Text: PrintColoredText(field, color)}
		cells = append(cells, &cell)
	}
	return cells
}

// GenerateCellsForBody - Generate Cells for body
func GenerateCellsForBody(values []string, size int) []*simpletable.Cell {
	var cells []*simpletable.Cell
	for i := 0; i < size; i++ {
		var cell simpletable.Cell
		if i == 0 {
			cell = simpletable.Cell{Align: simpletable.AlignLeft, Text: fmt.Sprintf(values[i][0:30])}
		} else {
			cell = simpletable.Cell{Align: simpletable.AlignLeft, Text: fmt.Sprintf(values[i])}
		}
		cells = append(cells, &cell)
	}
	return cells
}

// ShowTable - Show Table Contents
func ShowTable(fields []string, values []string) {
	table := simpletable.New()
	var cells []*simpletable.Cell

	for i := 0; i < len(fields); i++ {
		cells = GenerateCellsForHeader(fields, "yellow")
	}

	table.Header = &simpletable.Header{
		Cells: cells,
	}

	var cellsForBody []*simpletable.Cell
	for i := 0; i < len(fields); i++ {
		cellsForBody = GenerateCellsForBody(values, len(fields))
		table.Body.Cells = append(table.Body.Cells, cellsForBody)
	}

	table.SetStyle(simpletable.StyleDefault)
	fmt.Println(table.String())
}

// GetResourceMetada - Get struct information via reflection
func GetResourceMetada(resourceGroups interface{}) ([]string, []string) {
	return GetAllFieldsAndValues(resourceGroups)
}

// DisplayTable - Show Table Contents
func DisplayTable(resourceGroups interface{}) {
	table := simpletable.New()
	fields, values := GetAllFieldsAndValues(resourceGroups)
	var cells []*simpletable.Cell

	for i := 0; i < len(fields); i++ {
		cells = GenerateCellsForHeader(fields, "yellow")
	}

	table.Header = &simpletable.Header{
		Cells: cells,
	}

	var cellsForBody []*simpletable.Cell
	for i := 0; i < len(values); i++ {
		cellsForBody = GenerateCellsForBody(values, len(values))
		table.Body.Cells = append(table.Body.Cells, cellsForBody)
	}

	table.SetStyle(simpletable.StyleDefault)
	fmt.Println(table.String())
}

func CheckError(e error) {
	if e != nil {
		panic(e)
	}
}

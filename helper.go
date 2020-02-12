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

	"github.com/common-nighthawk/go-figure"
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

// ShowBanner - Show CLI Banner
func ShowBanner() {
	figure := figure.NewFigure("YASUKETE", "shadow", true)
	figure.Print()
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
func GenerateCellsForHeader(azureResource interface{}, color string) []*simpletable.Cell {
	var cells []*simpletable.Cell
	resType, size := GetTypeAttributes(azureResource)
	for i := 0; i < size; i++ {
		field := resType.Field(i).Name
		cell := simpletable.Cell{Align: simpletable.AlignLeft, Text: PrintColoredText(field, color)}
		cells = append(cells, &cell)
	}
	return cells
}

// GenerateCellsForBody - Generate Cells for body
func GenerateCellsForBody(values []string) []*simpletable.Cell {
	var cells []*simpletable.Cell
	for i := 0; i < len(values); i++ {
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

// DisplayTable - Show Table Contents
func DisplayTable(resourceGroups interface{}) {
	table := simpletable.New()
	s := reflect.ValueOf(resourceGroups)
	fmt.Println(s.Interface())

	cells := GenerateCellsForHeader(s.Type().Field(0), "yellow")
	table.Header = &simpletable.Header{
		Cells: cells,
	}

	//var cellsForBody []*simpletable.Cell
	// for _, row := range s.Index(0). {
	// 	f, v := GetAllFieldsAndValues(&row)
	// 	for i := 0; i < len(f); i++ {
	// 		cellsForBody = GenerateCellsForBody(v)
	// 		table.Body.Cells = append(table.Body.Cells, cellsForBody)
	// 	}
	// }

	table.SetStyle(simpletable.StyleDefault)
	fmt.Println(table.String())
}

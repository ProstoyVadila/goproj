package output

import (
	"fmt"
	"os"
	"strings"

	"github.com/elliotchance/orderedmap/v2"
	"github.com/fatih/color"
)

var infoMsg = color.New(color.FgMagenta, color.Bold)
var errMsg = color.New(color.FgRed)
var fieldShowMsg = color.New(color.FgHiWhite, color.Bold)
var valueShowMsg = color.New(color.FgCyan)

// Info prints a pretty colored info messages into output.
func Info(msg string, args ...interface{}) {
	if len(args) == 0 {
		infoMsg.Println(msg)
	} else {
		infoMsg.Printf(msg, args...)
	}
}

// Err prints a pretty colored error into output.
func Err(err error, msg ...string) {
	out := "error: " + err.Error()
	if len(msg) == 0 {
		errMsg.Println(out)
	} else {
		errMsg.Println(out + "\n" + msg[0])
	}
}

// Fatal prints a pretty colored error into output and exits with code 1.
func Fatal(err error, msg ...string) {
	Err(err, msg...)
	os.Exit(1)
}

// ShowString generates a string from orderedmap of model's fields and values in pretty readable way.
func ShowString(omap *orderedmap.OrderedMap[string, any], msg ...string) string {
	var builder strings.Builder
	var newLine = "\n"

	builder.WriteString(newLine)
	if len(msg) != 0 {
		builder.WriteString(infoMsg.Sprintln(msg[0]))
	}

	for el := omap.Front(); el != nil; el = el.Next() {
		line := fieldShowMsg.Sprintf(el.Key, valueShowMsg.Sprint(el.Value))
		builder.WriteString(line + "\n")
	}
	builder.WriteString(newLine)

	return fmt.Sprintln(builder.String())
}

// Show prints a string of model's fields and values in pretty readable way.
func Show(omap *orderedmap.OrderedMap[string, any], msg ...string) {
	fmt.Println(ShowString(omap, msg...))
}

package lab4

import (
	"fmt"
	"strings"
	"testing"

	engine "github.com/UniversalCorn/lab4/commands"
	parser "github.com/UniversalCorn/lab4/parser"
)

var cmds engine.Command

func BenchmarkParser(b *testing.B) {
	const baseLen = 5000
	for i := 0; i < 20; i++ {
		l := baseLen * (i + 1)

		input := "print " + strings.Repeat("smthsmthtoprintsmthtoprintsmthtoprint", l)

		b.Run(fmt.Sprintf("len=%d", l), func(b *testing.B) {
			cmds = parser.Parse(input)
		})
	}
}

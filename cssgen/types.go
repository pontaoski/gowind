package cssgen

import (
	"fmt"
	"strings"
	"unicode"

	"gowind/config"
)

type CSSLeaf struct {
	Class  string
	Styles map[string]string
}

type CSSDocument struct {
	Leaves     []CSSLeaf
	DarkLeaves []CSSLeaf
}

type Module interface {
	Generate(*config.Config) CSSDocument
}

func (c *CSSDocument) Consume(other *CSSDocument) {
	c.Leaves = append(c.Leaves, other.Leaves...)
	c.DarkLeaves = append(c.DarkLeaves, other.DarkLeaves...)
}

func (c *CSSDocument) String(cnf *config.Config) string {
	sb := strings.Builder{}

	write := func(s ...string) {
		for _, it := range s {
			sb.WriteString(it)
		}
	}
	writef := func(format string, a ...interface{}) {
		sb.WriteString(fmt.Sprintf(format, a...))
	}

	dumpLeaves := func(prefix string, leaves []CSSLeaf) {
		for _, leaf := range leaves {
			escape := func(s string) string {
				if unicode.IsNumber(rune(s[0])) {
					s = "\\" + s
				}

				s = strings.ReplaceAll(s, ":", "\\:")

				return s
			}
			write(".", escape(prefix+leaf.Class), " {\n")

			for k, v := range leaf.Styles {
				writef("\t%s: %s;\n", k, v)
			}

			write("}\n")
		}
	}

	dumpDarkLight := func(s string) {
		dumpLeaves(s, c.Leaves)
		write("@media (prefers-color-scheme: dark) {\n")
		dumpLeaves(s, c.DarkLeaves)
		write("}\n")
	}

	dumpDarkLight("")
	for name, size := range cnf.Screens.Screens {
		writef("@media (min-width: %s) {\n", size)
		dumpDarkLight(name + ":")
		writef("}\n")
	}

	return sb.String()
}

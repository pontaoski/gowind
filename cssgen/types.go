package cssgen

import (
	"fmt"
	"strings"

	"gowind/config"
)

type CSSLeaf struct {
	Selector string
	Rules    map[string]string
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

func (c *CSSDocument) String() string {
	sb := strings.Builder{}

	dumpLeaves := func(leaves []CSSLeaf) {
		for _, leaf := range leaves {
			sb.WriteString(".")
			sb.WriteString(leaf.Selector)
			sb.WriteString(" {\n")

			for k, v := range leaf.Rules {
				sb.WriteString(fmt.Sprintf("\t%s: %s;", k, v))
			}

			sb.WriteString("\n}\n")
		}
	}

	dumpLeaves(c.Leaves)
	sb.WriteString("@media (prefers-color-scheme: dark) {\n")
	dumpLeaves(c.DarkLeaves)
	sb.WriteString("}\n")

	return sb.String()
}

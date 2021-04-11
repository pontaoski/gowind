package colors

import (
	"fmt"
	"gowind/config"
	"gowind/cssgen"
)

type Module struct{}

func do(row *config.ColorRow, out *cssgen.CSSDocument, prop, classname string) {
	for name, shade := range row.Shades {
		out.Leaves = append(out.Leaves, cssgen.CSSLeaf{
			Selector: fmt.Sprintf("%s-%s-%s", classname, row.Name, name),
			Rules: map[string]string{
				prop: shade,
			},
		})
		if shade, ok := row.DarkVariants[name]; ok {
			out.DarkLeaves = append(out.DarkLeaves, cssgen.CSSLeaf{
				Selector: fmt.Sprintf("%s-%s-%s", classname, row.Name, name),
				Rules: map[string]string{
					prop: shade,
				},
			})
		}
		if name == "normal" {
			out.Leaves = append(out.Leaves, cssgen.CSSLeaf{
				Selector: fmt.Sprintf("%s-%s", classname, row.Name),
				Rules: map[string]string{
					prop: shade,
				},
			})
			if shade, ok := row.DarkVariants[name]; ok {
				out.DarkLeaves = append(out.DarkLeaves, cssgen.CSSLeaf{
					Selector: fmt.Sprintf("%s-%s", classname, row.Name),
					Rules: map[string]string{
						prop: shade,
					},
				})
			}
		}
	}
}

func (Module) Generate(c *config.Config) (d cssgen.CSSDocument) {
	for _, color := range c.Colors.Rows {
		do(color, &d, "color", "fg")
		do(color, &d, "background-color", "bg")
	}

	return
}

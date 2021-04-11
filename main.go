package main

import (
	"gowind/config"
	"gowind/cssgen"

	"gowind/mods/colors"
)

var mods = []cssgen.Module{
	colors.Module{},
}

func main() {
	config, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}
	doc := cssgen.CSSDocument{}
	for _, mod := range mods {
		d := mod.Generate(config)
		doc.Consume(&d)
	}
	print(doc.String())
}

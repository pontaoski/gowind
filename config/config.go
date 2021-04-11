package config

import "github.com/hashicorp/hcl/v2/hclsimple"

type ColorRow struct {
	Name         string            `hcl:"name,label"`
	Shades       map[string]string `hcl:"shades"`
	DarkVariants map[string]string `hcl:"dark,optional"`
}

type ColorsConfig struct {
	Rows []*ColorRow `hcl:"color,block"`
}

type Config struct {
	Colors ColorsConfig `hcl:"Colors,block"`
}

func LoadConfig() (*Config, error) {
	var config Config

	err := hclsimple.DecodeFile("gowind.hcl", nil, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}

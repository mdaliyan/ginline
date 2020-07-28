package config

import flags "github.com/spf13/pflag"


var (
	Address string
)

func ParseFlags() {

	flags.BoolP("verbose", "v", false, "verbose output")
	flags.StringVar(&Address, "listen", ":8080", "listen address")
	flags.PrintDefaults()
}


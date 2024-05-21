package command

import (
	"github.com/spf13/pflag"
)

var Mode string = "production"
var Profile string = ""

func init() {
	pflag.StringVar(&Mode, "mode", "production", "Run application in debug|production mode")
	pflag.StringVar(&Profile, "profile", "", "Run application with profile")
	pflag.Parse()
}
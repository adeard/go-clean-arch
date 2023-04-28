package debug

import (
	"os"

	"github.com/Code-Hex/dd/p"
)

func Dd(myVar ...interface{}) {
	p.P(myVar...)

	os.Exit(1)
}

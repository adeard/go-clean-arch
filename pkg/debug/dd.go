package debug

import (
	"github.com/Code-Hex/dd/p"
)

func Dd(myVar ...interface{}) {
	p.P(myVar...)
}

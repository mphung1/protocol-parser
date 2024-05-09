package protocols

import (
	"project/types"
)

func Dump (p types.IParser, _ types.IntOrString, _ types.IntOrString, _ types.IntOrString) {
	p.DumpStacks();
}
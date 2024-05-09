package protocols

import (
	"project/types"
)

func Push (p types.IParser, char types.IntOrString, val types.IntOrString, _ types.IntOrString) {
	stack := p.GetStack(char.(string), true);
	stack.Push(val.(int));
}
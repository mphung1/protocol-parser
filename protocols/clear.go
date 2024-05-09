package protocols

import (
	"project/types"
)

func Clear (p types.IParser, char types.IntOrString, _ types.IntOrString, _ types.IntOrString) {
	stack := p.GetStack(char.(string), false);

	if stack == nil {
		// Stack does not exist
		p.SetError(3);
		return;
	}
	
	stack.Clear();
}
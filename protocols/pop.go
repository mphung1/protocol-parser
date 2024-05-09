package protocols

import (
	"project/types"
	"project/utils"
)

func Pop (p types.IParser, char types.IntOrString, to types.IntOrString, _ types.IntOrString) {
	stack := p.GetStack(char.(string), false);

	if stack == nil {
		// Stack does not exist
		p.SetError(3);
		return;
	}

	_, ok := utils.SafeGet(p.GetArr(), to.(int));
	if !ok {
		// Out of bounds
		p.SetError(0);
		return;
	}

	if stack.IsEmpty() {
		p.SetError(2);
		return;
	}
	
	p.GetArr()[to.(int)] = stack.Pop();
}
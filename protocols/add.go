package protocols

import (
	"project/types"
	"project/utils"
)

func Add (p types.IParser, opera1 types.IntOrString, opera2 types.IntOrString, to types.IntOrString) {
	res := 0;

	// If the operand is a character, it is referencing a stack. It must pop the value from the stack.
	if p.IsString(opera1) {
		stack := p.GetStack(opera1.(string), false);
		if stack == nil {
			// Stack does not exist
			p.SetError(3);
			return;
		}

		res += stack.Pop();
	// Otherwise, it is a number
	} else {
		res += opera1.(int);
	}

	if p.IsString(opera2) {
		stack := p.GetStack(opera2.(string), false);
		if stack == nil {
			// Stack does not exist
			p.SetError(3);
			return;
		}

		res += stack.Pop();
	} else {
		res += opera2.(int);
	}

	_, ok := utils.SafeGet(p.GetArr(), to.(int));
	if !ok {
		// Out of bounds
		p.SetError(0);
		return;
	}
	
	p.GetArr()[to.(int)] = res;
}
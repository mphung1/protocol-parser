package protocols

import (
	"fmt"
	"project/types"
	"project/utils"
)

func Out (p types.IParser, from types.IntOrString, _ types.IntOrString, _ types.IntOrString) {
	// If from is a number, print the value at that index in the array
	if p.IsInt(from) {
		_, ok := utils.SafeGet(p.GetArr(), from.(int));
		if !ok {
			// Out of bounds
			p.SetError(0);
			return;
		}
		
		p.Print(fmt.Sprintf("%d", p.GetArr()[from.(int)]));
	// Otherwise, it is a character referencing a stack. Print the whole stack.
	} else {
		p.PrintStack(from.(string));
	}
}
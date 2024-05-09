package protocols

import (
	"project/types"
	"strings"
)

func Mode (p types.IParser, str types.IntOrString, _ types.IntOrString, _ types.IntOrString) {
	// First we are going to split the string into individuals chars,
	// because both could be enabled at the same time.
	modes := strings.Split(str.(string), "");

	for _, mode := range modes {
		// Toggle on each mode
		m := p.GetModeMap()[mode];
		
		if m == nil {
			// Mode does not exist
			p.SetError(5);
			return;
		}

		m();
	}
}
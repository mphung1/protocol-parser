package maps

import (
	"project/types"
)

// Pointer increase depending on protocol parameter count
func PointerIncreaseMap () types.PointerIncMap {
	pointer_map := make(types.PointerIncMap);
	
	pointer_map[11] = 2;
	// End protocol marks the main loop to stop due to negative pointer increase
	pointer_map[19] = -1000000000;
	pointer_map[20] = 4;
	pointer_map[21] = 4;
	pointer_map[30] = 3;
	pointer_map[31] = 3;
	pointer_map[32] = 2;
	pointer_map[33] = 1;
	pointer_map[40] = 2;

	return pointer_map;
}
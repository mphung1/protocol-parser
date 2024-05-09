package maps

import "project/types"

func ModeMap(p types.IParser) types.ModeMap {
	mode_map := make(types.ModeMap)

	mode_map["L"] = p.ToggleLogMode
	mode_map["R"] = p.ToggleRefMode

	return mode_map
}
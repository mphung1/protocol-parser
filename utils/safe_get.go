package utils

import "project/types"

// Safely gets an element from an array without going oob
func SafeGet(arr []types.IntOrString, index int) (types.IntOrString, bool) {
    if index >= 0 && index < len(arr) {
        return arr[index], true
    }

    return 0, false
}
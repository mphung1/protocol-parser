package maps

import (
	"project/exceptions"
	"project/types"
)

func ExceptionMap() types.ExceptionMap {
	exception_map := make(types.ExceptionMap)

	exception_map[0] = exceptions.IndexOutOfBoundsException.Error(exceptions.IndexOutOfBoundsException{})
	exception_map[1] = exceptions.IncompleteProtocolException.Error(exceptions.IncompleteProtocolException{})
	exception_map[2] = exceptions.EmptyStackException.Error(exceptions.EmptyStackException{})
	exception_map[3] = exceptions.NonExistentStackException.Error(exceptions.NonExistentStackException{})
	exception_map[4] = exceptions.InvalidCommandException.Error(exceptions.InvalidCommandException{})
	exception_map[5] = exceptions.InvalidModeException.Error(exceptions.InvalidModeException{})

	return exception_map
}
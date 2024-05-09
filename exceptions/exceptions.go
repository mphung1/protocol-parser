package exceptions

type IncompleteProtocolException struct{}

func (err IncompleteProtocolException) Error() string {
	return "IncompleteProtocolException"
}

type IndexOutOfBoundsException struct{}

func (err IndexOutOfBoundsException) Error() string {
	return "IndexOutOfBoundsException"
}

type InvalidCommandException struct{}

func (err InvalidCommandException) Error() string {
	return "InvalidCommandException"
}

type InvalidModeException struct{}

func (err InvalidModeException) Error() string {
	return "InvalidModeException"
}

type EmptyStackException struct{}

func (err EmptyStackException) Error() string {
	return "EmptyStackException"
}

type NonExistentStackException struct{}

func (err NonExistentStackException) Error() string {
	return "NonExistentStackException"
}

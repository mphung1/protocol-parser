package types

type Parser struct {
	// Array pointer
	Ptr 		int
	// Number array
	Arr 		[]IntOrString
	// If there was an error
	Error 		NotNilThenErr
	// Output
	Output 		[]string
	// Stack Cache
	StackCache 	StackCache
	// Log Line Counter
	LogLine 	int
	// Logging Mode
	LogMode 	bool
	// reference Mode
	RefMode 	bool
}

type IParser interface {
	DumpStacks() 
	GetArr() []IntOrString
	GetInstruction() (Instuction, int)
	GetModeMap() ModeMap
	GetOutput() []string
	GetPtr() int
	GetStack(char string, ispush bool) IStack
	GetStackCache() StackCache
	HandleProtoArg(arg IntOrString) IntOrString
	IsInt(s interface{}) bool
	IsString(s interface{}) bool
	Print(s string)
	PrintStack(char string)
	Parse()
	SetError(err int)
	ToggleLogMode()
	ToggleRefMode()
}

type ParserOutput struct {
	Output []string
}

type Stack struct {
	Arr []int
}

type IStack interface {
	Clear()
	GetStack() []int
	IsEmpty() bool
	Pop() int
	Push(val int)
}

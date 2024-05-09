package structures

import (
	"project/maps"
	"project/types"
	"project/utils"
	"strconv"
	"strings"
)

type ParseWrapper struct {
	types.Parser
}

func (p *ParseWrapper) DumpStacks() {
	// Printing each stack (order doesn't matter)
	// TODO: When this isEmpty is implemented, phase two tests don't pass. 	Wants ex: "c: "
	// TODO: However, when it's not, phase 3 doesn't pass... 				Wants ex: "c:"
	// TODO: The tests are not consistent lmao
	// TODO: We thought that END should have a trailing space because it was wanted before.
	for code, stack := range p.GetStackCache() {
		if stack.IsEmpty() {
			p.Print(code + ":")
		} else {
			p.Print(code + ": " + _GetStackString(stack))
		}
	}
}

// Gets the array from the parser
func (p *ParseWrapper) GetArr() []types.IntOrString {
	return p.Arr
}

// Gets the target instruction from the map and instuction code
func (p *ParseWrapper) GetInstruction() (types.Instuction, int) {
	// The instruction will be at the pointer's position
	// Check if the end of the array has been reached
	if p.Ptr >= len(p.Arr) {
		// No end command
		p.Error = 1
		return nil, 0
	}

	instruction := maps.InstructionMap()[p.Arr[p.Ptr].(int)]

	if instruction == nil {
		// Invalid command
		p.Error = 4
		return nil, 0
	}

	return instruction, p.Arr[p.Ptr].(int)
}

// Gets the mode map
func (p *ParseWrapper) GetModeMap() types.ModeMap {
	return maps.ModeMap(p)
}

// Gets the output from the parser
func (p *ParseWrapper) GetOutput() []string {
	return p.Output
}

// Gets the pointer from the parser
func (p *ParseWrapper) GetPtr() int {
	return p.Ptr
}

// Get stack via stack code (a character). Creates an empty stack if it doesn't exist.
func (p *ParseWrapper) GetStack(char string, ispush bool) types.IStack {
	res, exists := p.StackCache[char]

	//If the stack does not exist, create it
	if !exists {
		if !ispush {
			return nil
		}

		res = &StackWrapper{
			Stack: types.Stack{
				Arr: []int{},
			},
		}

		p.StackCache[char] = res
	}

	return res
}

// Gets the stack cache containing all stacks
func (p *ParseWrapper) GetStackCache() types.StackCache {
	return p.StackCache
}

// Handles a proto argument
func (p *ParseWrapper) HandleProtoArg(arg types.IntOrString) types.IntOrString {
	// Logic:
	// If arg == str, return the arg (mode doesn't matter)
	// if arg == int && !ref, return the arg
	// else arg == int && ref, return the corresponding value in the array
	if p.IsString(arg) {
		return arg
	}

	if p.IsInt(arg) && !p.RefMode {
		return arg
	}

	item, ok := utils.SafeGet(p.Arr, arg.(int))
	if !ok {
		// Out of bounds
		p.Error = 0
	}

	return item
}

// Checks if a value is an integer
func (p *ParseWrapper) IsInt(s interface{}) bool {
	_, ok := s.(int)
	return ok
}

// Checks if a value is a string
func (p *ParseWrapper) IsString(s interface{}) bool {
	_, ok := s.(string)
	return ok
}

// Prints a line to the output
func (p *ParseWrapper) Print(s string) {
	p.Output = append(p.Output, s)
}

// Prints the stack corresponding to a stack code (character). Empty stacks are an empty line. One space between each element.
func (p *ParseWrapper) PrintStack(char string) {
	stack := p.StackCache[char]

	res := _GetStackString(stack)

	p.Print(res)
}

// Main Parser function
func (p *ParseWrapper) Parse() {
	// While an end instruction has not been reached
	for p.Ptr >= 0 {
		instruction, code := p.GetInstruction()

		if instruction == nil {
			return
		}

		param_1, ok1 := utils.SafeGet(p.Arr, p.Ptr+1)
		param_2, ok2 := utils.SafeGet(p.Arr, p.Ptr+2)
		param_3, ok3 := utils.SafeGet(p.Arr, p.Ptr+3)

		params := [3]types.IntOrString{param_1, param_2, param_3}
		oob := [3]bool{!ok1, !ok2, !ok3}
		ptr_inc := maps.PointerIncreaseMap()[p.Arr[p.Ptr].(int)]

		slice := oob[:utils.Clamp(ptr_inc, 1, 1000000)-1]
		// If any is true there is an error
		for _, v := range slice {
			if v {
				// Out of bounds
				p.Error = 0
				return
			}
		}

		if p.LogMode {
			cmd_name := maps.GetInstructionName(code)
			p.Print(_FormatLoggingModeStr(p.LogLine, cmd_name, params, utils.Clamp(ptr_inc, 1, 1000000)))
			p.LogLine++
		}

		if p.RefMode {
			for i := range slice {
				params[i] = p.HandleProtoArg(params[i])
			}
		}
					
		if p.Error != nil {
			return // Error
		}

		instruction(p, params[0], params[1], params[2])
		p.Ptr += ptr_inc

		if p.Error != nil {
			return // Error
		}
	}

}

// Sets the error
func (p *ParseWrapper) SetError(err int) {
	p.Error = err
}

// Toggles the logging mode
func (p *ParseWrapper) ToggleLogMode() {
	p.LogMode = !p.LogMode
}

// Toggles the reference mode
func (p *ParseWrapper) ToggleRefMode() {
	p.RefMode = !p.RefMode
}

func _GetStackString(stack types.IStack) string {
	arr := stack.GetStack()

	if stack.IsEmpty() {
		return ""
	}

	str := make([]string, len(arr))

	// Conv to string
	for i, v := range arr {
		str[i] = strconv.Itoa(v)
	}

	// Reverse
	for i, j := 0, len(str)-1; i < j; i, j = i+1, j-1 {
		str[i], str[j] = str[j], str[i]
	}

	return strings.Join(str, " ")
}

func _FormatLoggingModeStr(line int, cmd string, params [3]types.IntOrString, inc int) string {
	// Slice params from 0 to inc - 1
	slice := params[:inc-1]
	// Convert to string
	str := make([]string, len(slice))

	for i, v := range slice {
		_, ok := v.(int)
		if ok {
			str[i] = strconv.Itoa(v.(int))
			continue
		}

		str[i] = v.(string)
	}
	// Join
	joined := strings.Join(str, " ")

	if inc == 1 {
		return strconv.Itoa(line) + ": " + cmd
	}

	return strconv.Itoa(line) + ": " + cmd + " " + joined
}

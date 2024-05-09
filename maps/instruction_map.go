package maps

import (
	"project/protocols"
	"project/types"
	"reflect"
	"runtime"
	"strings"
)

// Instruction to protocol mapping
func InstructionMap () types.InstuctionMap {
	instruct_map := make(types.InstuctionMap);
	
	instruct_map[11] = protocols.Out;
	instruct_map[19] = protocols.End;
	instruct_map[20] = protocols.Add;
	instruct_map[21] = protocols.Sub;
	instruct_map[30] = protocols.Push;
	instruct_map[31] = protocols.Pop;
	instruct_map[32] = protocols.Clear;
	instruct_map[33] = protocols.Dump;
	instruct_map[40] = protocols.Mode;

	return instruct_map;
}

// Use reflection to get the name of the instruction automatically
func GetInstructionName(code int) string {
	f := InstructionMap()[code]
	name := runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()
	return strings.ToUpper(strings.Split(name, ".")[1])
}
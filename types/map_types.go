package types

// func (*Parser, any, any, any): void
type Instuction func(IParser, IntOrString, IntOrString, IntOrString)

// {int: func (*Parser, Integer, Integer, Integer): void}
type InstuctionMap map[int]Instuction

type ExceptionMap map[int]string

// {int: int}
type PointerIncMap map[int]int

// {string: []int}
type StackCache map[string]IStack

type ModeMap map[string]func()
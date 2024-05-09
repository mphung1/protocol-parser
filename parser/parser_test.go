// DO NOT CHANGE THIS FILE
package parser

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPhase1(t *testing.T) {
	tests := map[string]struct {
		sInputValues   string
		expectedOutput []string
	}{
		`Test 1`: {
			sInputValues:   "20 20 5 14 21 7 18 14 11 14 11 4 19 12 32 19",
			expectedOutput: []string{"-11", "21"},
		},
		`Test 2`: {
			sInputValues:   "20 10 11 4 20 40 9 0 11 0 19",
			expectedOutput: []string{"31"},
		},
		`Test 3`: {
			sInputValues:   "19",
			expectedOutput: nil,
		},
		`Test 4`: {
			sInputValues:   "11 1 11 2 11 3 11 4 11 5 19",
			expectedOutput: []string{"1", "11", "2", "11", "3"},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			inputValues := strings.Split(test.sInputValues, " ")
			po := &ParserOutput{}

			Parse(inputValues, po)

			require.Len(t, po.Output, len(test.expectedOutput))
			for index, value := range test.expectedOutput {
				assert.Equal(t, value, po.Output[index])
			}
		})
	}
}

func TestPhase2(t *testing.T) {
	tests := map[string]struct {
		sInputValues   string
		expectedOutput []string
	}{
		`Test 1`: {
			sInputValues:   "30 a 3 30 a 5 30 b 4 30 b 9 31 a 22 31 b 22 31 a 23 20 -1 -1 0 11 0 19",
			expectedOutput: []string{"12"},
		},
		`Test 2`: {
			sInputValues:   "30 a 3 30 b 5 30 b 4 20 b a 0 11 0 19",
			expectedOutput: []string{"7"},
		},
		`Test 3`: {
			sInputValues:   "30 a 5 30 a 8 30 a 128 11 a 19",
			expectedOutput: []string{"128 8 5"},
		},
		`Test 4`: {
			sInputValues:   "30 x 1 30 x 32 21 x x 0 11 0 19",
			expectedOutput: []string{"31"},
		},
		`Test 5`: {
			sInputValues:   "30 j 14 21 j 5 0 11 0 19",
			expectedOutput: []string{"9"},
		},
		`Test 6`: {
			sInputValues:   "30 j 14 21 5 j 0 11 0 19",
			expectedOutput: []string{"-9"},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			inputValues := strings.Split(test.sInputValues, " ")
			po := &ParserOutput{}

			Parse(inputValues, po)

			require.Len(t, po.Output, len(test.expectedOutput))
			for index, value := range test.expectedOutput {
				assert.Equal(t, value, po.Output[index])
			}
		})
	}

	// Special handling for dump test because I really messed up in defining the spec
	// and this isn't as deterministic as I wish it was
	dumpValues := "30 a 3 30 a 2 30 a 1 30 b -1 30 b 13 30 c 15 32 c 30 d 31 33 32 b 33 19"

	inputValues := strings.Split(dumpValues, " ")
	po := &ParserOutput{}

	Parse(inputValues, po)

	dumpExpectedOutput := []string{
		"a: 1 2 3",
		"b: 13 -1",
		"c:",
		"d: 31",
		"a: 1 2 3",
		"b:",
		"c:",
		"d: 31",
	}

	assert.Len(t, po.Output, len(dumpExpectedOutput))

	for _, expected := range dumpExpectedOutput {
		assert.Contains(t, po.Output, expected)
	}
}

func TestPhase3(t *testing.T) {
	tests := map[string]struct {
		sInputValues   string
		expectedOutput []string
	}{
		`Test 1`: {
			sInputValues:   "40 R 20 0 2 3 11 3 21 7 4 3 11 3 19",
			expectedOutput: []string{"60", "1"},
		},
		`Test 2`: {
			sInputValues:   "40 R 30 20 21 30 22 23 30 24 25 20 20 20 26 11 26 33 19 0 a 12 a 31 a 17 0",
			expectedOutput: []string{"48", "a: 12"},
		},
		`Test 3`: {
			sInputValues: "40 L 20 20 5 16 21 7 18 16 11 16 11 6 19 12 32 19",
			expectedOutput: []string{
				"1: ADD 20 5 16",
				"2: SUB 7 18 16",
				"3: OUT 16",
				"-11",
				"4: OUT 6",
				"21",
				"5: END",
			},
		},
		`Test 4`: {
			sInputValues: "40 L 30 a 5 30 a 8 30 a 800 11 a 19",
			expectedOutput: []string{
				"1: PUSH a 5",
				"2: PUSH a 8",
				"3: PUSH a 800",
				"4: OUT a",
				"800 8 5",
				"5: END",
			},
		},
		`Test 6`: {
			sInputValues: "40 RL 20 0 2 3 11 3 21 7 4 3 11 3 19",
			expectedOutput: []string{
				"1: ADD 0 2 3",
				"2: OUT 3",
				"60",
				"3: SUB 7 4 3",
				"4: OUT 3",
				"1",
				"5: END",
			},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			inputValues := strings.Split(test.sInputValues, " ")
			po := &ParserOutput{}

			Parse(inputValues, po)

			require.Len(t, po.Output, len(test.expectedOutput))
			for index, value := range test.expectedOutput {
				assert.Equal(t, value, po.Output[index])
			}
		})
	}

	// Special handling for dump test because I really messed up in defining the spec
	// and this isn't as deterministic as I wish it was
	dumpValues := "40 L 30 a 3 30 a 2 30 a 1 30 b -1 30 b 13 30 c 15 32 c 30 d 31 33 32 b 33 19"

	inputValues := strings.Split(dumpValues, " ")
	po := &ParserOutput{}

	Parse(inputValues, po)

	dumpExpectedOutput := []string{
		"1: PUSH a 3",
		"2: PUSH a 2",
		"3: PUSH a 1",
		"4: PUSH b -1",
		"5: PUSH b 13",
		"6: PUSH c 15",
		"7: CLEAR c",
		"8: PUSH d 31",
		"9: DUMP",
		"a: 1 2 3",
		"b: 13 -1",
		"c:",
		"d: 31",
		"10: CLEAR b",
		"11: DUMP",
		"a: 1 2 3",
		"b:",
		"c:",
		"d: 31",
		"12: END",
	}

	assert.Len(t, po.Output, len(dumpExpectedOutput))

	for _, expected := range dumpExpectedOutput {
		assert.Contains(t, po.Output, expected)
	}

}

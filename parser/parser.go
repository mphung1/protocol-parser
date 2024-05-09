package parser

import (
	"project/maps"
	"project/structures"
	"project/types"
	"project/utils"
)

func Parse (values []string, output *ParserOutput) types.NotNilThenErr {
	pw := structures.ParseWrapper {
		Parser:	types.Parser{
			Ptr:		0, 
			Arr:		utils.InputToArr(values),
			Error:		nil,
			Output:		[]string{},
			StackCache:	types.StackCache{},
			LogLine:	1,
			LogMode:	false,
			RefMode:	false,
		},
	};

	pw.Parse();
	//pw.PrintArr();

	output.Output = pw.GetOutput();

	if pw.Error != nil {
		return maps.ExceptionMap()[pw.Error.(int)];
	}

	return nil;
}
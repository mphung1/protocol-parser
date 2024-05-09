package utils

import (
	"project/types"
	"strconv"
)

func InputToArr(arr []string) []types.IntOrString {
	// Create another array to hold the strings as integers
	int_arr := make([]types.IntOrString, 0, len(arr));
	// Map the strings to integers where necessary
	for i := 0; i < len(arr); i++ {
		num, err := strconv.Atoi(arr[i]);
		if err == nil {
			int_arr = append(int_arr, num);
		} else {
			int_arr = append(int_arr, arr[i]);
		}
	}

	return int_arr;
}
package oasa_sync_utils

import (
	"fmt"
	"strconv"
	"strings"
)

func StrToInt64(numStr string) int64 {
	var result int64 = 0
	trimmedStr := strings.TrimSpace(numStr)
	var res, err = strconv.ParseInt(trimmedStr, 0, 64)
	if err != nil {
		fmt.Printf("An Error occurred in conversation from Str to Int64!", err.Error())
	} else {
		result = res
	}
	return result
}

func StrToInt16(numStr string) int16 {
	var result int16 = 0
	trimmedStr := strings.TrimSpace(numStr)
	var res, err = strconv.ParseInt(trimmedStr, 0, 16)
	if err != nil {
		fmt.Printf("An Error occurred in conversation from Str to Int64!", err.Error())
	} else {
		result = int16(res)
	}
	return result
}

func StrToInt(numStr string) int {
	var result int = 0
	trimmedStr := strings.TrimSpace(numStr)
	var res, err = strconv.ParseInt(trimmedStr, 0, 0)
	if err != nil {
		fmt.Printf("An Error occurred in conversation from Str to Int64!", err.Error())
	} else {
		result = int(res)
	}
	return result
}

func StrToFloat(numStr string) float32 {
	var result float32 = 0
	if numStr != "" {
		trimmedStr := strings.TrimSpace(numStr)
		//trimmedStr = strings.Replace(trimmedStr, ".", ",", -1)
		var res, err = strconv.ParseFloat(trimmedStr, 32)
		if err != nil {
			fmt.Printf("An Error occurred in conversation from Str to Int64!", err.Error())
		} else {
			result = float32(res)
		}
	}

	return result
}

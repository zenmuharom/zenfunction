package function

import (
	"fmt"
	"strings"
)

func escapedCommas(str string) (newStr string) {
	newStr = strings.ReplaceAll(str, ",", `\,`)
	return
}

func unEscapedCommas(str string) (newStr string) {
	newStr = strings.ReplaceAll(str, `\,`, ",")
	return
}

func splitWithEscapedCommas(str string) []string {
	var splitStrings []string
	var currentString string
	escaped := false

	for _, char := range str {
		if char == '\\' && !escaped {
			escaped = true
			continue
		}

		if char == ',' && !escaped {
			// fmt.Println(fmt.Sprintf("char:%v | escaped:%v", char, escaped))
			fmt.Println("masuk ada koma")
			splitStrings = append(splitStrings, currentString)
			currentString = ""
			continue
		}

		currentString += string(char)
		escaped = false
		fmt.Println(currentString)
	}

	// splitStrings = append(splitStrings, strings.TrimSpace(currentString))
	fmt.Println("currentString:")
	fmt.Println(currentString)
	splitStrings = append(splitStrings, currentString)
	fmt.Println("splitstring:")
	fmt.Println(splitStrings)
	return splitStrings
}

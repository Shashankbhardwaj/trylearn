package main

import (
	"fmt"
	"strings"
)

func main() {
	var (
		testCases     int
		inputStacks   []string
		inputStack    string
		lenStack      int
		maneuverCount int
	)
	fmt.Println("Enter the Test Cases \n")
	fmt.Scanf("%d", &testCases)
	fmt.Printf("\nEnter %d input Strings\n", testCases)
	for i := 0; i < testCases; i++ {
		fmt.Scanf("%s", &inputStack)
		lenStack = len(inputStack)
		stringStack := strings.Split(inputStack, "")
		fmt.Println("\n", stringStack, lenStack, "\n")
		maneuverCount = maneuver(stringStack, lenStack-1, 0)

		inputStacks = append(inputStacks, inputStack)
	}
	fmt.Println(inputStacks, maneuverCount)
}

func maneuver(inputStack []string, downFaceCount, maneuverCount int) int {

	fmt.Println("inside maneuver \n", inputStack, downFaceCount, maneuverCount, "\n")
	fmt.Println("before calling findDownfacecount \n", inputStack, downFaceCount, "\n")
	downFaceCount = findNextDownFaceCount(inputStack, downFaceCount)
	fmt.Println("after calling findDownfacecount \n", inputStack, downFaceCount, maneuverCount, "\n")
	if downFaceCount >= 0 {
		maneuverCount = maneuverCount + 1
		for j := downFaceCount; j >= 0; j-- {
			fmt.Println("before calling revertFace\n", inputStack, j, "\n")
			inputStack = revertFace(inputStack, j)
			fmt.Println("after calling revertFace\n", inputStack, j, "\n")
		}
		fmt.Println("before calling maneuver \n", inputStack, downFaceCount, maneuverCount, "\n")
		maneuverCount = maneuver(inputStack, downFaceCount-1, maneuverCount)
	}
	fmt.Printf("ManeuverCount %d", maneuverCount)
	return maneuverCount

}

func findNextDownFaceCount(inputString []string, previousDownFaceCount int) int {

	for j := previousDownFaceCount; j >= 0; j-- {
		if inputString[j] == "-" {
			return j
		}
	}

	return -1
}

func revertFace(inputString []string, count int) []string {
	if inputString[count] == "-" {
		inputString[count] = "+"
	} else {
		inputString[count] = "-"
	}
	return inputString
}

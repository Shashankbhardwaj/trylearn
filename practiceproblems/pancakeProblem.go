package practiceproblems

import (
	"fmt"
)

func PancakeProblem() {
	var (
		testCases  int
		inputStack string
		// lenStack      int
		maneuverCount int
		inputStacks   []string
	)
	// fmt.Println("Enter the Test Cases \n")
	fmt.Scanf("%d", &testCases)
	// fmt.Printf("\nEnter %d input Strings\n", testCases)
	for i := 0; i < testCases; i++ {
		fmt.Scanf("%s", &inputStack)
		inputStacks = append(inputStacks, inputStack)
	}

	for i := 0; i < testCases; i++ {
		maneuverCount = 0
		// lenStack = 0
		// lenStack = len(inputStacks[i])
		// stringStack := strings.Split(inputStacks[i], "")
		// fmt.Println("\n", stringStack, lenStack, "\n")
		// maneuverCount = maneuver(stringStack, lenStack-1, 0)
		maneuverCount = maneuverWitoutChange(inputStacks[i])
		fmt.Printf("\nCase #%d: %d\n", i+1, maneuverCount)
	}
}

func maneuver(inputStack []string, downFaceCount, maneuverCount int) int {

	// fmt.Println("inside maneuver \n", inputStack, downFaceCount, maneuverCount, "\n")
	// fmt.Println("before calling findDownfacecount \n", inputStack, downFaceCount, "\n")
	downFaceCount = findNextDownFaceCount(inputStack, downFaceCount)
	// fmt.Println("after calling findDownfacecount \n", inputStack, downFaceCount, maneuverCount, "\n")
	if downFaceCount >= 0 {
		maneuverCount = maneuverCount + 1
		for j := downFaceCount; j >= 0; j-- {
			// fmt.Println("before calling revertFace\n", inputStack, j, "\n")
			inputStack = revertFace(inputStack, j)
			// fmt.Println("after calling revertFace\n", inputStack, j, "\n")
		}
		// fmt.Println("before calling maneuver \n", inputStack, downFaceCount, maneuverCount, "\n")
		maneuverCount = maneuver(inputStack, downFaceCount-1, maneuverCount)
	}
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

func maneuverWitoutChange(inputStack string) int {

	maneuverCount := 0
	firstFaceDown := 0
	for j := len(inputStack) - 1; j >= 0; j-- {
		if inputStack[j] == '-' {
			firstFaceDown = j
			break
		}
	}

	for i := firstFaceDown; i >= 0; i-- {
		if i < len(inputStack)-1 {
			if inputStack[i] != inputStack[i+1] {
				maneuverCount++
			}
		} else if i == len(inputStack)-1 {
			maneuverCount++
		}
	}

	return maneuverCount
}

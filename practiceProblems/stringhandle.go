package practiceProblems

import (
	"fmt"
)

func StringHandle() {

	stringArr := []string{"4", "11", "3", "5", "-10", "3", "-5", "-8", "9", "12", "-13", "-14", "15", "20", "-6"}
	var result string
	var maxElement, secondMaxElement, val string
	for _, stringNum := range stringArr {
		fmt.Println(maxElement, stringNum)
		val = FindMaxElement(maxElement, stringNum)
		if val == stringNum {
			secondMaxElement = maxElement
			maxElement = val
		} else {
			secondMaxElement = FindMaxElement(secondMaxElement, stringNum)
		}

	}

	if secondMaxElement == "" {
		result = "-1"
	}

	result = secondMaxElement

	fmt.Println("Second Maximum Element", result)

}

func FindMaxElement(var1, var2 string) string {

	if var1 == "" {
		return var2
	} else if var2 == "" {
		return var1
	}

	isnegative := false
	if var1[0] == '-' && var2[0] != '-' {
		return var2
	} else if var1[0] != '-' && var2[0] == '-' {
		return var1
	} else if var1[0] == '-' && var2[0] == '-' {
		isnegative = true
	}

	var1length := len(var1)
	var2length := len(var2)

	if var1length > var2length {
		if isnegative {
			return var2
		}
		return var1
	}
	if var2length > var1length {
		if isnegative {
			return var1
		}
		return var2
	} else if var1length == var2length {
		for i := 0; i < var1length; i++ {
			if var1[i] > var2[i] {
				if isnegative {
					return var2
				}
				return var1
			} else if var1[i] < var2[i] {
				if isnegative {
					return var1
				}
				return var2
			}
		}
	}
	return var1
}

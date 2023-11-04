package util

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	minSum = 9
	maxSum = 19
)

func GeneratePassword(input UserInput) string {
	serviceName := input.Service
	secretKey := input.SecretKey
	secretNumbers := input.SecretNumbers

	firstSymbol := getSymbol(serviceName, 0)
	lastSymbol := getSymbol(serviceName, len(serviceName)-1)

	num1, num2, num3 := extractNumsFromSecretKey(firstSymbol, lastSymbol)
	secretNum1, secretNum2 := extractNumsFromSecretNumbers(secretNumbers)
	secretNum3 := convertNumsToSecretNum(num1, num2, num3)
	encodeNumbers := secretNum1 + secretNum2 + secretNum3

	return fmt.Sprintf("%v%v_%s_%d", strings.ToUpper(firstSymbol), strings.ToLower(lastSymbol), secretKey, encodeNumbers)
}

func convertNumsToSecretNum(num1, num2, num3 int) int {
	return num1*100 + num2*10 + num3
}

func extractNumsFromSecretKey(firstSymbol, lastSymbol string) (int, int, int) {
	num1 := getNumBySymbol(strings.ToUpper(firstSymbol))
	num2 := getNumBySymbol(strings.ToUpper(lastSymbol))
	num3 := calcLastNumberSequence(num1, num2)
	return num1, num2, num3
}

func extractNumsFromSecretNumbers(secretNumbers string) (int, int) {
	middle := len(secretNumbers) / 2
	secretNum1, _ := strconv.Atoi(secretNumbers[:middle])
	secretNum2, _ := strconv.Atoi(secretNumbers[middle:])
	return secretNum1, secretNum2
}

func getNumBySymbol(symbol string) int {
	var symbolToNums = map[string]int{
		"A": 2, "B": 2, "C": 2,
		"D": 3, "E": 3, "F": 3,
		"G": 4, "H": 4, "I": 4,
		"J": 5, "K": 5, "L": 5,
		"M": 6, "N": 6, "O": 6,
		"P": 7, "Q": 7, "R": 7, "S": 7,
		"T": 8, "U": 8, "V": 8,
		"W": 9, "X": 9, "Y": 9, "Z": 9,
	}
	return symbolToNums[strings.ToUpper(symbol)]
}

func getSymbol(str string, index int) string {
	return string(str[index])
}

func calcLastNumberSequence(num1, num2 int) (num3 int) {
	if num1+num2 > minSum {
		num3 = maxSum - (num1 + num2)
	} else {
		num3 = minSum - (num1 + num2)
	}
	return
}

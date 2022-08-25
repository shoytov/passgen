package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

var numbersArray = []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
var specialSymbolsArray = []string{"~", "!", "&", "?", ",", ":", "-", "_", "<", ">", "'", "%", "#", "$"}
var symbols = []string{"q", "w", "e", "r", "t", "y", "u", "i", "o", "p", "a", "s", "d", "f", "g", "h", "j", "k", "l",
	"z", "x", "c", "v", "b", "n", "m"}
var randomCondition = []int{2, 3, 5}

func shuffleRandomCondition() {
	rand.Seed(time.Now().UnixNano())

	for i := len(randomCondition) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		randomCondition[i], randomCondition[j] = randomCondition[j], randomCondition[i]
	}
}

func removeFromSlice(arrayPointer *[]string, array []string, index int) {
	*arrayPointer = append(array[:index], array[index+1:]...)
}

func findIndexOfElementInSlice(array []string, el string) int {
	result := -1
	for i := 0; i < len(array); i++ {
		if array[i] == el {
			result = i
			break
		}
	}
	return result
}

func clearByNumber(number string) {
	if number == "1" {
		index := findIndexOfElementInSlice(numbersArray, number)
		if index > -1 {
			removeFromSlice(&numbersArray, numbersArray, index)
		}

		index = findIndexOfElementInSlice(symbols, "l")
		if index > -1 {
			removeFromSlice(&symbols, symbols, index)
		}
	}

	if number == "0" {
		index := findIndexOfElementInSlice(numbersArray, number)
		if index > -1 {
			removeFromSlice(&numbersArray, numbersArray, index)
		}

		index = findIndexOfElementInSlice(symbols, "o")
		if index > -1 {
			removeFromSlice(&symbols, symbols, index)
		}
	}
}

func clearByLetter(letter string) {
	if letter == "l" || letter == "i" {
		index := findIndexOfElementInSlice(numbersArray, "1")
		if index > -1 {
			removeFromSlice(&numbersArray, numbersArray, index)
		}

		index = findIndexOfElementInSlice(symbols, "l")
		if index > -1 {
			removeFromSlice(&symbols, symbols, index)
		}

		index = findIndexOfElementInSlice(symbols, "i")
		if index > -1 {
			removeFromSlice(&symbols, symbols, index)
		}
	}

	if letter == "o" {
		index := findIndexOfElementInSlice(numbersArray, "0")
		if index > -1 {
			removeFromSlice(&numbersArray, numbersArray, index)
		}

		index = findIndexOfElementInSlice(symbols, letter)
		if index > -1 {
			removeFromSlice(&symbols, symbols, index)
		}
	}
}

func main() {
	shuffleRandomCondition()
	resPass := ""
	defaultLength := 8
	rand.Seed(time.Now().Unix())

	if len(os.Args) > 1 {
		i, err := strconv.Atoi(os.Args[1])

		if err == nil {
			defaultLength = i
		}
	}

	for i := 0; i < defaultLength; i++ {
		if i%randomCondition[0] == 0 {
			symbol := symbols[rand.Intn(len(symbols))]
			resPass += symbol
			clearByLetter(symbol)
			continue
		}
		if i%randomCondition[1] == 0 {
			resPass += specialSymbolsArray[rand.Intn(len(specialSymbolsArray))]
			continue
		}
		if i%randomCondition[2] == 0 {
			symbol := numbersArray[rand.Intn(len(numbersArray))]
			resPass += symbol
			clearByNumber(symbol)
		} else {
			symbol := symbols[rand.Intn(len(symbols))]
			resPass += strings.ToUpper(symbol)
			clearByLetter(symbol)
		}
	}
	fmt.Println(resPass)
}

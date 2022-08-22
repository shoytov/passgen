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

func main() {
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
		if i%2 == 0 {
			resPass += symbols[rand.Intn(len(symbols))]
			continue
		}
		if i%3 == 0 {
			resPass += specialSymbolsArray[rand.Intn(len(specialSymbolsArray))]
			continue
		}
		if i%5 == 0 {
			resPass += numbersArray[rand.Intn(len(numbersArray))]
		} else {
			resPass += strings.ToUpper(symbols[rand.Intn(len(symbols))])
		}
	}
	fmt.Println(resPass)
}

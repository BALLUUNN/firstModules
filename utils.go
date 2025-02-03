package utils

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

const NumberOfSymbol int = 3
//GetFloat reads a float-point number from the keyboard.
//It returns the number read and any error encountered.
func GetFloat() (float64, error) {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		return 0, nil
	}
	input = strings.TrimSpace(input)
	grate, err := strconv.ParseFloat(input, 64)
	if err != nil {
		return 0, err
	}
	return grate, nil
}

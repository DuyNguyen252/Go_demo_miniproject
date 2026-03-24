package utils

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

// Nhập string
func InputString(label string) string {
	fmt.Print(label)
	text, _ := reader.ReadString('\n')
	return strings.TrimSpace(text)
}

// Nhập int
func InputInt(label string) int {
	for {
		input := InputString(label)
		num, err := strconv.Atoi(input)
		if err == nil {
			return num
		}
		fmt.Println("Nhập số không hợp lệ, thử lại!")
	}
}

// Nhập float
func InputFloat(label string) float64 {
	for {
		input := InputString(label)
		num, err := strconv.ParseFloat(input, 64)
		if err == nil {
			return num
		}
		fmt.Println("Nhập số không hợp lệ, thử lại!")
	}
}

package external

import (
	"fmt"
	"strings"
)

func PrintBox(message string) {
	padding := 2
	content := fmt.Sprintf("%s%s%s", strings.Repeat(" ", padding), message, strings.Repeat(" ", padding)) // Add padding around the message
	boxWidth := len(content)

	// Top and bottom borders with '-'
	topBottomBorder := "+" + strings.Repeat("-", boxWidth) + "+"

	// Print the top border
	fmt.Println(topBottomBorder)
	// Print the message with side borders
	fmt.Printf("|%s|\n", content)
	// Print the bottom border
	fmt.Println(topBottomBorder)
}

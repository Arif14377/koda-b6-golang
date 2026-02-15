package views

import "fmt"

func Input(prompt string) int {
	var input int
	fmt.Printf("\n%s", prompt)
	fmt.Scanf("%d", &input)
	return input
}

package views

import "fmt"

func Input(prompt string) int {
	var input int
	fmt.Printf("%s", prompt)
	fmt.Scanf("%d\n", &input)
	return input
}

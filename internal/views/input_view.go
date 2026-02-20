package views

import "fmt"

func InputAngka(prompt string) int {
	var input int
	fmt.Printf("%s", prompt)
	fmt.Scanf("%d\n", &input)
	return input
}

func InputString(prompt string) string {
	var input string
	fmt.Printf("%s", prompt)
	fmt.Scanf("%s\n", &input)
	return input
}

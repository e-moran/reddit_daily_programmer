package main

import (
	"fmt"
)

func main() {
	var first_word, second_word string;

	fmt.Println("Please enter the first word for the word funnel: ")
	fmt.Scanf("%s", &first_word)
	fmt.Println("Please enter the second word for the word funnel: ")
	fmt.Scanf("%s", &second_word)

	fmt.Println(word_funnel(first_word, second_word))
}

func word_funnel(first string, second string) bool {
	mismatch_count := 0

	if(len(first) != len(second) +1) {
		return false
	}

	for i := 0; i < len(second); i++ {
		if (first[:i+mismatch_count] + first[i+1+mismatch_count:]) == second {
			mismatch_count++
		}
	}
	return mismatch_count == 1
}
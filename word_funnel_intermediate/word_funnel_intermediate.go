package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var word string
	word_list := read_word_list()

	fmt.Println("Please enter the first word for the word funnel: ")
	fmt.Scanf("Maximum Word Funnel Chain: %s", &word)

	fmt.Println(word_funnel_count(word, word_list, 1))
}

func read_word_list() []string {
	file, err := os.Open("wordlist.txt")
	if err != nil {
		return nil
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}

func word_funnel_count(word string, list []string, count int) (curr int) {
	curr = count
	length := count
	if len(word) > 1 {
		for _, element := range list {
			is, funnel := is_word_funnel(word, element)

			if is {
				length = word_funnel_count(funnel, list, 1+count)
				if length > curr {
					curr = length
				}
			}
		}
	}

	return
}

func is_word_funnel(first string, second string) (is bool, funnel string) {
	mismatch_count := 0
	is = false
	funnel = ""

	if(len(first) != len(second) +1) {
		return
	}

	for i := 0; i < len(second); i++ {
		if (first[:i+mismatch_count] + first[i+1+mismatch_count:]) == second {
			mismatch_count++
			funnel = second
		}
	}

	//Special case for if last character removed
	if (first[:len(second)+mismatch_count]) == second {
		mismatch_count++
		funnel = second
	}

	is = mismatch_count == 1
	return
}
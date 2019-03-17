package main

import (
	"bufio"
	"fmt"
	"os"
)

//dup2

func main() {
	counts := make(map[string]map[string]int)
	files := os.Args[1:]
	for _, arg := range files {
		f, err := os.Open(arg)
		if err != nil {
			fmt.Println(err)
			continue
		}
		counts[string(arg)] = make(map[string]int)
		CountLines(f, counts[string(arg)])
		f.Close()
	}

	for filename, moji := range counts {
		// if n > 1 {
		// 	fmt.Printf("%d %s\n", n, line)
		// }
		fmt.Printf("%s\n", filename)
		for key, value := range moji {
			// fmt.Printf("%d %s\n",value,key)
			if value > 1 {
				fmt.Printf("%d %s\n", value, key)
			}
		}
		fmt.Println()
	}
}

func CountLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
}

//dup1

// func main() {
// 	counts := make(map[string]int)
// 	input := bufio.NewScanner(os.Stdin)
// 	for input.Scan() {
// 		counts[input.Text()]++
// 	}
// 	for line, count := range counts {
// 		if count > 1 {
// 			fmt.Printf("%d\t%s\n", count, line)
// 		}
// 	}
// }

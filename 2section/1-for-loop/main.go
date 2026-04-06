package main

import "fmt"

func main() {

	// for -- only way to loop

	// C-Style loop
	for i := 1; i < 10; i++ {
		// fmt.Println(i)
	}

	// While-style
	k := 3
	for k > 0 {
		fmt.Println(k)
		k--
	}

	fmt.Println(" ----- Infinite Loop -----")
	count := 0
	for {
		fmt.Println("counter:,", count)
		count++
		if count >= 100 {
			break
		}
	}

	fmt.Println(" ----- skipping -----")
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println(i)
	}

	fmt.Println(" ------- array ---------")
	items := [3]string{"Go", "Python", "Java"}
	for _, v := range items {
		fmt.Println(v)
	}

}

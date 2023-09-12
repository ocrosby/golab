// You can edit this code!
// Click here and start typing.
package main

import "fmt"

func main() {
	// make new
	// new - only allocates - no init of memory
	// make - allocates and initializes - non zeroed storage

	score := make(map[string]int)

	score["omar"] = 89
	score["josh"] = 34
	score["ron"] = 24
	score["sam"] = 12
	score["vicky"] = 44

	fmt.Println(score)

	getRonScore := score["ron"]

	fmt.Println(getRonScore)

	delete(score, "vicky")

	fmt.Println(score)

	for k, v := range score {
		fmt.Printf("Score of %v is %v\n", k, v)
	}
}

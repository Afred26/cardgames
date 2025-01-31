package card

import "fmt"

func Example_rank() {

	r1 := Ace
	r2 := Rank(11)

	fmt.Println(r1)
	fmt.Println(r2)

	// Output:
	//A
	//K
}

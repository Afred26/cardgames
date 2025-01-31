package card

import "fmt"

func Example_suit() {

	s1 := Spades
	s2 := Suit(3)

	fmt.Println(s1)
	fmt.Println(s2)

	// Output:
	//♠
	//♣
}

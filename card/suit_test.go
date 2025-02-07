package card

import "fmt"

func ExampleSuit_String() {
	s1 := Spades
	s2 := Suit(4)

	fmt.Println(s1)
	fmt.Println(s2)

	// Output:
	// ♠
	// invalid suit
}

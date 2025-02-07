package card

import "fmt"

func ExampleCard_String() {
	c1 := New(Ace, Spades)
	c2 := New(King, Hearts)
	c3 := New(Ten, Diamonds)

	fmt.Println(c1)
	fmt.Println(c2)
	fmt.Println(c3)

	// Output:
	//┌───────┐
	//│A      │
	//│       │
	//│   ♠   │
	//│       │
	//│      A│
	//└───────┘
	//┌───────┐
	//│K      │
	//│       │
	//│   ♥   │
	//│       │
	//│      K│
	//└───────┘
	//┌───────┐
	//│10     │
	//│       │
	//│   ♦   │
	//│       │
	//│     10│
	//└───────┘

}

func ExampleCard_Matches() {
	c1 := New(Ace, Spades)
	c2 := New(Queen, Hearts)
	c3 := New(Ten, Hearts)

	fmt.Println(c1.Matches(c2))
	fmt.Println(c2.Matches(c3))
	fmt.Println(c3.Matches(c2))
	fmt.Println(c3.Matches(c1))

	// Output:
	// false
	// true
	// true
	// false
}

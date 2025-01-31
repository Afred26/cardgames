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
	//│     A │
	//└───────┘
	//┌───────┐
	//│K      │
	//│       │
	//│   ♥   │
	//│       │
	//│     K │
	//└───────┘
	//┌───────┐
	//│10     │
	//│       │
	//│   ♦   │
	//│       │
	//│     10│
	//└───────┘

}

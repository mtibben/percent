package percent

import (
	"fmt"
)

func Example() {
	percentEncoded := Escape("a.b/c d", "/. ")
	fmt.Println(percentEncoded)
	fmt.Println(Unescape(percentEncoded))
	// Output: a%2Eb%2Fc%20d
	// a.b/c d
}

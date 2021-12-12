// Have the query setup as a string.

// All of the paramters will simply be in objects here as string fields
// Then I can call the strings in the query in the other main file.

package dbseed

import (
	"fmt"
)

var mb = [...]string {
	"blurb1",
	"I think this is a bit better.",
}

func TestPrint() {
	fmt.Println("The package dbseed has been called and printed")
}
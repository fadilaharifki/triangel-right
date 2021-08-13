// result
// *
// **
// ***
// ****
// *****

package main

import "fmt"

func main() {
	result := ""

	for i := 0; i < 5; i++ {
		for j := 0; j < 1; j++ {
			result += "*"
		}
		fmt.Println(result)
	}
}

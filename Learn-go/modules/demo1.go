package b

import (
	"fmt"
	_ "fmt"
)

func PassByVlaueAndReference() {
	a := make(map[string]string)
	b := "Pramod"
	c := "Bhavith"
	a["first"] = b
	a["second"] = c
	fmt.Println(a, "Before swap", b, c)
	private1(&b, &c)
	fmt.Println(a, "After swap", b, c)
	primes := [6]int{2, 3, 5, 7, 11, 13}

	var s []int = primes[1:4]
	fmt.Println(s)

}

func private1(a *string, b *string) {
	var e = *a
	*a = *b
	*b = e
	fmt.Println("Public1")
}

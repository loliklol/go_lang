package main

import "fmt"

func main() {
	var hello = "Hello \n\t"
	println(hello)
	var rawBinary byte = '\x27'
	println(rawBinary)

	test := "aseweq"
	//test2 := 13
	test3 := `!!!!!fsd
	  	qw	 	qw	f!!"`
	new := test + test3

	println(new)

	const (
		t1 = iota + 9
		t2
		_
		t3
		t4   = iota
		size = 5
	)

	var a1 [3]int
	var a2 [2 * size]bool
	a3 := [...]int{1, 2, 3}
	a1[0] = 7
	fmt.Print(a1)
	fmt.Print(a2)
	fmt.Print(a3)
	a7:=0
	a8:=0

    a4 :=[...]{a7, a8}
	/*const b = 3

	const c int = 4
	*/
	println(t1, t2, t3, t4)
``
	/*

		struct {

			type=[int,float,string]
			intval
			floatval
			stringval
			objref
		}



	*/
	//приведение типов prinln(i+ int(i))

}

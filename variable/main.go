package main

import "fmt"

var x int = 1

func main() {
	// Membuat variable
	var myName string = "Noobee id"

	var yourName string
	yourName = "David"

	myAge := 42

	fmt.Println(myName)
	fmt.Println(yourName)
	fmt.Println(myAge)
	fmt.Println(x)
	// variable tanpa nilai awal
	var x string
	var y int
	var z bool

	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)

	// Deklarasi multiple variable
	var a, b, c, d = 1, 2, 3, "halo"
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)

	//deklarasi multiple variable dalam blok
	var (
		ab string = "nobee id"
		ac        = "Id"
		ad int    = 170
	)
	fmt.Println(ab)
	fmt.Println(ac)
	fmt.Println(ad)
}

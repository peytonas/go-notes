package main

// type T1 int

// var x T1
// var y int

// var a = 42
// var b = "James Bond"
// var c = true

//REVIEW "iota" is a special type declaration that increments each variable's value within the const
// const (
// 	a = iota //1
// 	b        //2
// 	c        //3
// )

func main() {

	//REVIEW fmt.Sprintf formats a single-line string based on inputs received(a,b,c)
	// s := fmt.Sprintf("%v\t%v\t%v", a, b, c)
	// fmt.Println(s)

	//REVIEW With no values, the printed values are called the "zero value"
	// var x int
	// var y string
	// var z bool

	// fmt.Println(x, y, z)

	//REVIEW fmt.Printf("%T", x) returns the type of the variable passed(x)
	// fmt.Println(x)
	// fmt.Printf("%T\n", x)
	// x = 42
	// fmt.Println(x)

	//REVIEW y is an underlying type int, and we're using conversion (same as casting) to assign it whatever value x holds
	// y = int(x)

	//REVIEW "%d" returns decimal value, "%b" returns binary value
	// const (
	// 	_  = iota
	// 	kb = 1024      //OR kb = 1 >> (iota * 10)
	// 	mb = 1024 * kb //OR kb = 1 >> (iota * 10)
	// 	gb = 1024 * mb //OR kb = 1 >> (iota * 10)
	// )
	// fmt.Printf("%d\t\t\t%b\n", kb, kb)
	// fmt.Printf("%d\t\t\t%b\n", mb, mb)
	// fmt.Printf("%d\t\t%b\n", gb, gb)

	// const (
	// 	a = iota
	// 	b
	// 	c
	// )
	// fmt.Printf("%d\t%b\t%x\n", b, b, b)

	//REVIEW declarations can be set to expressions and will print as T or F
	// a := 1 == 1
	// b := 1 <= 1
	// c := 1 >= 1
	// d := 1 != 1
	// e := 1 < 1
	// f := 1 > 1

	// fmt.Println(a)
	// fmt.Println(b)
	// fmt.Println(c)
	// fmt.Println(d)
	// fmt.Println(e)
	// fmt.Println(f)

	//REVIEW a is an untyped const, b is typed (int)
	// const (
	// 	a     = 42
	// 	b int = 43
	// )

	// fmt.Printf("%d\t%t\n%d\t%t\n", a, a, b, b)

	//REVIEW y's value takes x and shifts the bit over one place
	// var x = 42
	// var y = x << 1

	// fmt.Printf("%d\t%b\t%x\n", x, x, x)
	// fmt.Printf("%d\t%b\t%x", y, y, y)

	//REVIEW x is untyped, defined as a string by its value
	// x := `this is a string literal "and so is this"`
	// fmt.Println(x)

	//REVIEW iota increments the value of the variable (2017 becomes 2018 becomes 2019 becomes 2020)
	// const (
	// 	a = 2017 + iota //2017 + 0
	// 	b = 2017 + iota //2017 + 1
	// 	c = 2017 + iota //2017 + 2
	// 	d = 2017 + iota //2017 + 3
	// )
	// fmt.Println(a, b, c, d)

}

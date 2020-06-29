package main

// type Person struct {
// 	first string
// 	last  string
// 	age   int
// }

//REVIEW function takes in a param of type person and passes them the method speak
// func (p Person) speak() {
// 	fmt.Println("I am", p.first, p.last, "and I am", p.age, "years old.")
// }

//REVIEW foo takes in a series of ints, and returns an int, loops through/adds together the variadic slice, returns total
// func foo(nums ...int) int {
// 	total := 0
// 	for _, num := range nums {
// 		total += num
// 	}
// 	fmt.Println(total)
// 	return total
// }

//REVIEW same as foo except bar takes in a slice of ints, and returns an int, loops through/adds together the slice, returns total
// func bar(ints []int) int {
// 	total := 0
// 	for _, num := range ints {
// 		total += num
// 	}
// 	fmt.Println(total)
// 	return total
// }

// func foo() {
// 	fmt.Println("Hello World")
// }

// func bar() {
// 	fmt.Println("Goodbye World")
// }

//REVIEW Go fizzbuzz
// func fizzBuzz() {
// 	for i := 1; i <= 100; i++ {
// 		if i%15 == 0 {
// 			fmt.Printf("%v\tFizzBuzz\n", i)
// 		} else if i%5 == 0 {
// 			fmt.Printf("%v\tBuzz\n", i)
// 		} else if i%3 == 0 {
// 			fmt.Printf("%v\tFizz\n", i)
// 		} else {
// 			fmt.Printf("%v\tNA\n", i)
// 		}
// 	}
// }

//REVIEW "iota" is a special type declaration that increments each variable's value within the const
// const (
// 	a = iota //1
// 	b        //2
// 	c        //3
// )
// var x = 0
// var g = func() {
// 	x++
// 	fmt.Println("testing\n", x)
// }

func main() {

	// f := func() {
	// 	for i := 0; i < 3; i++ {
	// 		fmt.Println(i)
	// 	}
	// }
	// f()
	// g()
	// fmt.Printf("%T\n", f)
	// fmt.Printf("%T\n", g)

	//REVIEW intseq calls an anonymous func and returns an int; the anonymous func also returns an int ("i"), and provides closure for the returned func and variable i;
	// func() {
	// 	for i := 0; i <= 100; i++ {
	// 		fmt.Println(i)
	// 	}
	// }() //this second set of parens calls the anonymous function "func()"

	// p1 := Person{
	// 	first: "Peyton",
	// 	last:  "Sonnefeld",
	// 	age:   30,
	// }
	// p1.speak()

	//REVIEW defer forces the called function to wait til everything else has run (until main ends in this case), and then runs foo.
	// defer foo()
	// bar()
	// nums := []int{1, 2, 3, 4, 5}
	// foo(nums...)
	// bar(nums)

	// fizzBuzz()

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

	//REVIEW "for" works like a while loop, everything else is exactly as usual with looping
	// for x := 33; x <= 122; x++ {
	// 	fmt.Printf("%v\t%#U\n", x, x)
	// }

	// sum := 1
	// for sum < 1000 {
	// 	sum += sum
	// }
	// fmt.Println(sum)

	//REVIEW with no condition, loop runs until hitting a break
	// for {
	//       fmt.Println("loop")
	//       break
	// 	}

	//REVIEW You can also continue to the next iteration of the loop (won't continue until conditional evaluates to true)
	// for n := 0; n <= 5; n++ {
	// 	if n%2 == 0 {
	// 		continue
	// 	}
	// 	fmt.Println(n)
	// }

	// for i := 0; i < 100; i++ {
	// 	if i%2 == 0 {
	// 		fmt.Println(i)
	// 	}
	// }

	//REVIEW switch statements work exactly the same
	// switch {
	// case false:
	// 	fmt.Println("this should not print")
	// case 2 == 4:
	// 	fmt.Println("this should not print")
	// case 3 == 3:
	// 	fmt.Println("3 == 3")
	// 	fallthrough //REVIEW with fallthrough even when a case evaluates as true it will still go to the next case
	// case 4 == 4:
	// 	fmt.Println("4 == 4")
	// 	fallthrough
	// case 5 == 4:
	// 	fmt.Println("this should not print")
	// 	fallthrough
	// case 7 == 9:
	// 	fmt.Println("this should not print")
	// 	fallthrough
	// case 8 == 8:
	// 	fmt.Println("8 == 8")
	// default:
	// fmt.Println("default")
	// }

	// i := 65
	// for i <= 90 {
	// 	for j := 0; j < 3; j++ {
	// 		fmt.Printf("%#U\n", i)
	// 	}
	// 	i++
	// }

	// i := 2020
	// for i >= 1990 {
	// 	fmt.Println(i)
	// 	i--
	// }

	// i := 2020
	// for {
	// 	if i < 1990 {
	// 		break
	// 	}
	// 	fmt.Println(i)
	// 	i--
	// }

	// i := 10
	// for i <= 100 {
	// 	fmt.Println(i % 4)
	// 	i++
	// }

	// for i := 0; i <= 10; i++ {
	// 	if i%2 == 0 {
	// 		fmt.Println("even")
	// 	} else if i%2 != 0 {
	// 		fmt.Println("odd")
	// 	}
	// }

	// favSport := "swim"
	// switch favSport {
	// case "skiing":
	// 	fmt.Println("mountains")
	// case "hockey":
	// 	fmt.Println("HIT HIM")
	// default:
	// 	fmt.Println("I hate sports")
	// }

	//REVIEW "i" is the index where we are, "v" is the value at index "i" from array "x" (i.e. for each index "i" within array "x", do something)
	// x := [5]int{1, 2, 3, 4, 5}
	// for i, v := range x {
	// 	fmt.Println(i, v)
	// }
	// fmt.Printf("%T\n", x)

	//REVIEW slices work exactly the same as arrays, just don't specify length when value is declared
	// x := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	// for i, v := range x {
	// 	fmt.Println(i, v)
	// }
	// fmt.Printf("%T\n", x)

	//REVIEW to create new slices from an array, declare new slice variable and define the range you want from within the array("x")
	// x := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	// var y []int = x[0:5]
	// var z []int = x[5:10]
	// var a []int = x[2:7]
	// var b []int = x[1:6]
	// fmt.Printf("%v\n%v\n%v\n%v\n", y, z, a, b)
	// fmt.Printf("%T\n", x)

	//REVIEW append redefines a previously declared array ("x") with a new value pushed onto its end (52)
	// x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	// x = append(x, 52)
	// fmt.Println(x)

	//REVIEW append values from "x" up to position 3, AND values from position 6 to the end, then re-map array "x"
	// x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	// x = append(x[:3], x[6:]...)
	// fmt.Println(x)

	//REVIEW "make" creates a slice of type string(1st arg) with a length of 5(2nd arg) and a capacity of 5(3rd arg)
	// x := make([]string, 5, 5)
	// x = []string{`Alabama`, `Washington`, `California`, `Hawaii`, `Florida`} //this is a composite literal
	// fmt.Println(x)
	// fmt.Println(len(x))
	// fmt.Println(cap(x))

	// for i := 0; i < len(x); i++ {
	// 	fmt.Println(i, x[i])
	// }

	//REVIEW this creates a map where each key is a string, and each value for that key is a string
	// x := map[string][]string{
	// 	"bond_james":      []string{"Martinis", "Women", "Shaken, not stirred"},
	// 	"no_dr":           []string{`being evil`, `ice cream`, `sunsets`},
	// 	"moneypenny_miss": []string{`James Bond`, `Literature`, `Computer Science`},
	// }

	//REVIEW for each key("k") within each value("v") within map "x", print the key; then, for each value within each key, print the value.
	// for k, v := range x {
	// 	fmt.Println("This is for", k)
	// 	for i, v2 := range v {
	// 		fmt.Println("\t", i, v2)
	// 	}
	// }

	//REVIEW this is exactly the same as above except line 273 adds a new record to the map by declaring the key, and then defining its values as a slice of strings, with the values in that slice
	// x := map[string][]string{
	// 	"bond_james":      []string{"Martinis", "Women", "Shaken, not stirred"},
	// 	"no_dr":           []string{`being evil`, `ice cream`, `sunsets`},
	// 	"moneypenny_miss": []string{`James Bond`, `Literature`, `Computer Science`},
	// }

	// x["Q"] = []string{`Tech`, `Scotch`, `Cars`} // adds record "Q" to map x

	// for k, v := range x {
	// 	fmt.Println("This is for", k)
	// 	for i, v2 := range v {
	// 		fmt.Println("\t", i, v2)
	// 	}
	// }

	//REVIEW also exactly the same as above, but line 290 deletes a record from a map
	// x := map[string][]string{
	// 	"bond_james":      []string{"Martinis", "Women", "Shaken, not stirred"},
	// 	"no_dr":           []string{`being evil`, `ice cream`, `sunsets`},
	// 	"moneypenny_miss": []string{`James Bond`, `Literature`, `Computer Science`},
	// }

	// x["Q"] = []string{`Tech`, `Scotch`, `Cars`}

	// delete(x, `no_dr`) //deletes from map "x"(arg 1) the specified record(arg 2)

	// for k, v := range x {
	// 	fmt.Println("This is for", k)
	// 	for i, v2 := range v {
	// 		fmt.Println("\t", i, v2)
	// 	}
	// }

	//REVIEW defines type Person with its values
	// type Person struct {
	// 	firstName    string
	// 	lastName     string
	// 	faveIceCream []string
	// }

	//REVIEW creates types of Person
	// person1 := Person{"Peyton", "Sonnefeld", []string{"Half Baked", "Strawberry", "Vanilla"}}
	// person2 := Person{"Chelsea", "Dunn", []string{"Vanilla", "Chocolate"}}

	//REVIEW creates a map containing type Person, and declares the values (key=lastName, value=person1/person2)
	// m := map[string]Person{
	// 	person1.lastName: person1,
	// 	person2.lastName: person2,
	// }

	// fmt.Println(m)

	//REVIEW for each type in "m", print that type and its value
	// for _, v := range m {
	// 	fmt.Println(v.firstName, v.lastName)
	// 	//REVIEW for each value in faveIceCream, print the index and its value
	// 	for i, val := range v.faveIceCream {
	// 		fmt.Println(i, val)
	// 	}
	// 	fmt.Println("---------")
	// }

	//REVIEW declares Vehicle struct as an underlying struct of type Sedan and Truck
	// type Vehicle struct {
	// 	doors int
	// 	color string
	// }

	// type Truck struct {
	// 	Vehicle
	// 	fourWheel bool
	// }

	// type Sedan struct {
	// 	Vehicle
	// 	luxury bool
	// }

	// f150 := Truck{
	// 	Vehicle: Vehicle{
	// 		doors: 2,
	// 		color: "black"},
	// 	fourWheel: true}
	// taurus := Sedan{
	// 	Vehicle: Vehicle{
	// 		doors: 4,
	// 		color: "blue"},
	// 	luxury: true}

	// fmt.Println(f150)
	// fmt.Println(taurus)
	// fmt.Println(f150.color)
	// fmt.Println(taurus.doors)

	//REVIEW anonymous struct declared, same as structs above without a declared type; struct is given values, then types of the anonymous struct are declared
	// s := struct {
	// 	name      string
	// 	friends   map[string]int
	// 	favDrinks []string
	// }{
	// 	name: "Peyton",
	// 	friends: map[string]int{
	// 		"Chelsea": 444,
	// 		"Tom":     555,
	// 	},
	// 	favDrinks: []string{
	// 		"Monster",
	// 		"Water",
	// 	},
	// }

	// fmt.Println(s.name)
	// fmt.Println(s.friends)

}

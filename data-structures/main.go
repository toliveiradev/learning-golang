package main

import "fmt"

// creating struct methods
type Rectangle struct {
	length  float64
	breadth float64
}

func (r *Rectangle) area() float64 {
	return r.length * r.breadth
}

func (r *Rectangle) SetLength(length float64) {
	r.length = length
}

func (r *Rectangle) SetBreadth(breadth float64) {
	r.breadth = breadth
}

func main() {
	/*
		-- Arrays --
	**/
	// creating and assigning values to an array without var keyword
	studentsAge := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// creating a nested array
	nestedArray := [3][5]int{
		{1, 2, 3, 4, 5},
		{6, 7, 8, 9, 10},
		{11, 12, 13, 14, 15},
	}
	fmt.Println("studentsAge", studentsAge) // [1 2 3 4 5 6 7 8 9 10]
	fmt.Println("nestedArray", nestedArray) // [[1 2 3 4 5] [6 7 8 9 10] [11 12 13 14 15]]

	// accessing array values with their indexes
	fmt.Println("studentsAge[0]", studentsAge[0]) // 1
	fmt.Println("studentsAge[1]", studentsAge[1]) // 2
	fmt.Println("studentsAge[9]", studentsAge[9]) // 10

	// using a for loop to access an array
	for i := 0; i < len(studentsAge); i++ {
		fmt.Println("studentsAge[i]", studentsAge[i])
	}

	// using range to access an array
	for index, value := range studentsAge {
		fmt.Printf("index: %v, value: %v\n", index, value)
	}

	// modifying array values with their indexes
	studentsAge[0] = 35
	studentsAge[3] = 17
	studentsAge[8] = 20

	fmt.Println("studentsAge modified", studentsAge) // [35 2 3 17 5 6 7 8 20 10]

	// creating and getting the lenght of an array with a lenght of 10
	var arrayOfIntegers [10]int
	fmt.Println("len(arrayOfIntegers)", len(arrayOfIntegers)) // 10

	// creating and getting the lenght of an array with a lenght of 7
	var arrayOfStrings [7]string
	fmt.Println("len(arrayOfStrings)", len(arrayOfStrings)) // 7

	// creating and getting the lenght of an array with a lenght of 20
	var arrayOfBooleans [20]bool
	fmt.Println("len(arrayOfBooleans)", len(arrayOfBooleans)) // 20

	/*
		-- Slices --
	**/
	// creating slices from arrays
	fiveStudentrsAge := studentsAge[0:5]
	fmt.Println("fiveStudentrsAge", fiveStudentrsAge) // [35 2 3 17 5]
	threeStudentsAge := studentsAge[3:6]
	fmt.Println("threeStudentsAge", threeStudentsAge) // [17 5 6]

	// creating slices with make specifying lenght
	sliceOfIntegers := make([]int, 5)       // [0 0 0 0 0]
	sliceOfBooleans := make([]bool, 3)      // [false false false]
	sliceOfStrings := make([]string, 5, 10) // [     ]
	fmt.Println("sliceOfIntegers", sliceOfIntegers)
	fmt.Println("sliceOfBooleans", sliceOfBooleans)
	fmt.Println("sliceOfStrings", sliceOfStrings)

	// creating and assigning values to a slice without var keyword
	taskRemaining := []string{"eat", "sleep", "code", "repeat"}
	fmt.Println("taskRemaining", taskRemaining) // [eat sleep code repeat]

	// creating a nested slice
	nestedSlice := [][]int{
		{1},
		{2, 3},
		{4, 5, 6},
		{7, 8, 9, 10},
	}
	fmt.Println("nestedSlice", nestedSlice) // [[1] [2 3] [4 5 6] [7 8 9 10]]

	// creating a slice from literals
	sliceFromLiterals := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// accessing slice values with their indexes
	firstInteger := sliceFromLiterals[0]        // 1
	secondInteger := sliceFromLiterals[1]       // 2
	lastInteger := sliceFromLiterals[9]         // 10
	fmt.Println("firstInteger", firstInteger)   // 1
	fmt.Println("secondInteger", secondInteger) // 2
	fmt.Println("lastInteger", lastInteger)     // 10

	// using a for loop to access a slice
	for i := 0; i < len(sliceFromLiterals); i++ {
		fmt.Println("sliceFromLiterals[i]", sliceFromLiterals[i])
	}

	// using range to access a slice
	for index, value := range sliceFromLiterals {
		fmt.Printf("index: %v, value: %v\n", index, value)
	}

	// creating and getting the lenght of a slice
	sliceOfIntegers2 := make([]int, 5)                          // [0 0 0 0 0]
	fmt.Println("len(sliceOfIntegers2)", len(sliceOfIntegers2)) // 5

	// creating and getting the capacity of a slice
	sliceOfIntegers3 := make([]int, 5, 10)                      // [0 0 0 0 0]
	fmt.Println("cap(sliceOfIntegers3)", cap(sliceOfIntegers3)) // 10

	// creating a slice from literals
	sliceOfIntegers4 := []int{1, 2, 3}

	// using append to add a single value to the slice
	sliceOfIntegers4 = append(sliceOfIntegers4, 4)
	fmt.Println(sliceOfIntegers) // [1 2 3 4]

	// using append to add multiple values to the slice
	sliceOfIntegers4 = append(sliceOfIntegers4, 5, 6, 7)
	fmt.Println(sliceOfIntegers4) // [1 2 3 4 5 6 7]

	// using append to add a slice to a slice
	anotherSlice := []int{8, 9, 10}
	sliceOfIntegers = append(sliceOfIntegers, anotherSlice...)
	fmt.Println(sliceOfIntegers) // [1 2 3 4 5 6 7 8 9 10]

	/*
		-- Maps --
	**/
	// creating a string -> int map
	studentsAgeMap := make(map[string]int)
	fmt.Println("studentsAgeMap", studentsAgeMap) // map[]

	// creating a map from literals
	studentsAgeMap2 := map[string]int{
		"john":  32,
		"peter": 28,
		"jane":  45,
		"joe":   19,
		"linda": 20,
	}
	fmt.Println("studentsAgeMap2", studentsAgeMap2) // map[jane:45 joe:19 john:32 linda:20 peter:28]

	// creating nested maps
	studentResults := map[string]map[string]int{
		"john":  {"math": 78, "english": 80},
		"peter": {"math": 45, "english": 56},
	}
	fmt.Println("studentResults", studentResults)                             // map[john:map[english:80 math:78] peter:map[english:56 math:45]]
	fmt.Println("studentResults[john]", studentResults["john"])               // map[english:80 math:78]
	fmt.Println("studentResults[john][math]", studentResults["john"]["math"]) // 78

	// accessing values in the map
	fmt.Println("studentsAgeMap2[john]", studentsAgeMap2["john"])   // 32
	fmt.Println("studentsAgeMap2[peter]", studentsAgeMap2["peter"]) // 28
	fmt.Println("studentsAgeMap2[jane]", studentsAgeMap2["jane"])   // 45
	fmt.Println("studentsAgeMap2[joe]", studentsAgeMap2["joe"])     // 19
	fmt.Println("studentsAgeMap2[linda]", studentsAgeMap2["linda"]) // 20

	// two-value assignment to check if the key exists
	element, ok := studentsAgeMap2["john"]
	fmt.Println(element, ok) // 32 true

	// two-value assignment to check if the key does not exist
	element2, ok2 := studentsAgeMap2["james"]
	fmt.Println(element2, ok2) // 0 false

	// updating values in the map
	studentsAgeMap2["john"] = 35
	fmt.Println("studentsAgeMap2[john]", studentsAgeMap2["john"]) // 35

	// deleting keys from the map
	delete(studentsAgeMap2, "john")
	delete(studentsAgeMap2, "jane")
	fmt.Println("studentsAgeMap2", studentsAgeMap2) // map[joe:19 linda:20 peter:28]

	/*
		-- Structs --
	**/
	// creating a struct instance with var
	var myRectangle Rectangle
	fmt.Println("myRectangle", myRectangle) // {0 0}

	// creating an empty struct instance
	myRectangle = Rectangle{}
	fmt.Println("myRectangle", myRectangle) // {0 0}

	// creating a struct instance specifying values
	myRectangle = Rectangle{10.5, 5.5}
	fmt.Println("myRectangle", myRectangle) // {10.5 5.5}

	// creating a struct instance specifying fields and values
	myRectangle = Rectangle{length: 10.5, breadth: 5.5}
	fmt.Println("myRectangle", myRectangle) // {10.5 5.5}

	// we can also omit struct fields during their instantiation
	myRectangle = Rectangle{breadth: 5.5}
	fmt.Println("myRectangle", myRectangle) // {0 5.5}

	// creating an array and slice of structs
	arrayofRectangles := [5]Rectangle{
		{8.5, 6.5},
		{10.5, 5.5},
		{12.5, 7.5},
		{14.5, 4.5},
		{16.5, 8.5},
	}
	fmt.Println("arrayofRectangles", arrayofRectangles) // [{8.5 6.5} {10.5 5.5} {12.5 7.5} {14.5 4.5} {16.5 8.5}]

	sliceOfRectangles := []Rectangle{
		{8.5, 6.5},
		{10.5, 5.5},
		{12.5, 7.5},
		{14.5, 4.5},
		{16.5, 8.5},
	}
	fmt.Println("sliceOfRectangles", sliceOfRectangles) // [{8.5 6.5} {10.5 5.5} {12.5 7.5} {14.5 4.5} {16.5 8.5}]

	// creating a pointer struct instance
	myRectangle2 := &Rectangle{length: 10, breadth: 5}
	fmt.Println(myRectangle2, *myRectangle2) // &{10 5} {10 5}

	// creating a struct instance with new
	myRectangle3 := new(Rectangle)
	fmt.Println(myRectangle3, *myRectangle3) // &{0 0} {0 0}

	// creating a struct instance specifying fields and values
	myRectangle4 := Rectangle{length: 10, breadth: 5}

	// accessing the values in struct fields
	fmt.Println("myRectangle4.length", myRectangle4.length)   // 10
	fmt.Println("myRectangle4.breadth", myRectangle4.breadth) // 5

	myRectangle4.length = 15
	myRectangle4.breadth = 7
	fmt.Println("myRectangle4", myRectangle4) // {15 7}

	// creating a nested struct
	type address struct {
		houseNumber int
		streetName  string
		city        string
		state       string
		country     string
	}

	type Person struct {
		firstName   string
		lastName    string
		homeAddress address
	}

	// creating an instance of a nested struct
	person := Person{
		firstName: "John",
		lastName:  "Doe",
		homeAddress: address{
			houseNumber: 123,
			streetName:  "Main St",
			city:        "Cleveland",
			state:       "Ohio",
			country:     "USA",
		},
	}

	fmt.Println(("person.firstName"), person.firstName)             // John
	fmt.Println("person.homeAddress.city", person.homeAddress.city) // Cleveland

	// creating a struct anonymously
	circle := struct {
		radius float64
		color  string
	}{
		radius: 10,
		color:  "red",
	}

	fmt.Println("circle", circle)             // {10 red}
	fmt.Println("circle.color", circle.color) // red

	// creating a struct instance
	myRectangle5 := Rectangle{length: 5, breadth: 10}
	fmt.Println("Area of rectangle is: ", myRectangle5.area())

	// creating a struct instance
	myRectangle6 := Rectangle{length: 20, breadth: 10}
	fmt.Println("myRectangle6", myRectangle6) // {20 10}

	myRectangle6.SetLength(30)
	myRectangle6.SetBreadth(15)
	fmt.Println("myRectangle6", myRectangle6) // {30 15}
}

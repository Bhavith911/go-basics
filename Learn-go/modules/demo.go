package b

import "fmt"

func private() {
	fmt.Println("private")
}

func Public() {
	fmt.Println("Public")
}

func Declarations() {

	// Explicit type declaration
	var x int = 10
	fmt.Println("x:", x)

	// Type inference
	var y = 20
	fmt.Println("y:", y)

	// Short variable declaration (only inside functions)
	z := 30
	fmt.Println("z:", z)
}

func Statements() {
	// If statement
	x := 10
	if x > 5 {
		fmt.Println("x is greater than 5")
	}

	// For loop
	for i := 0; i < 3; i++ {
		fmt.Println(i)
	}

	// For As while loop
	i := 0
	for i < 5 {
		fmt.Println(i)
		i++
	}

	// infinite For loop
	for {
		fmt.Println(i)
		i++
		if i >= 5 {
			break // Break out of the loop when i reaches 5
		}
	}

	// For with range
	nums := []int{10, 20, 30, 40}
	for index, value := range nums {
		fmt.Printf("Index: %d, Value: %d\n", index, value)
	}

	fruits := map[string]string{"a": "Apple", "b": "Banana", "c": "Cherry"}
	for key, value := range fruits {
		fmt.Printf("Key: %s, Value: %s\n", key, value)
	}

	for _, value := range nums {
		fmt.Println(value)
	}

	// Only index
	for index := range nums {
		fmt.Println(index)
	}

	// for loop with labels

outerLoop:
	for i := 1; i <= 3; i++ {
	innerForLoop:
		for j := 1; j <= 3; j++ {
			if i*j > 4 {
				break outerLoop // Breaks the outer loop
				break innerForLoop
			}
			fmt.Printf("%d * %d = %d\n", i, j, i*j)
		}
	}

	// Switch statement
	day := "Monday"
	switch day {
	case "Monday":
		fmt.Println("Start of the week!")
	case "Friday":
		fmt.Println("End of the week!")
	default:
		fmt.Println("Midweek day!")
	}
}

func SliceMapArray() {
	m := map[string]int{"key1": 1, "key2": 2}
	modifyMap(m)
	fmt.Println("Modified map:", m)

	// Array (PBV)
	arr := [3]int{1, 2, 3}
	modifyArray(arr)
	fmt.Println("Array after modification:", arr)

	// Slice (PBR)
	slice := []int{1, 2, 3}
	modifySlice(slice)
	fmt.Println("Slice after modification:", slice)
}

func modifyMap(m map[string]int) {
	m["key2"] = 100 // Modifies the original map
}

func modifyArray(arr [3]int) {
	arr[0] = 100 // Does not modify the original array
}

func modifySlice(s []int) {
	s[0] = 100 // Modifies the original slice
}

func Structures() {
	// Define a struct
	type Person struct {
		Name string
		Age  int
	}

	// Create a new Person
	person := Person{Name: "John", Age: 25}
	fmt.Println("Person:", person)

	// Access fields
	fmt.Println("Name:", person.Name)
	fmt.Println("Age:", person.Age)
}

// Function
func add(x, y int) int {
	return x + y
}

// Method for struct Person
type Person struct {
	Name string
	Age  int
}

func (p Person) Greet() {
	fmt.Printf("Hello, my name is %s!\n", p.Name)
}

func FunctionMethods() {
	// Call function
	fmt.Println("Sum:", add(5, 3))

	// Call method
	person := Person{Name: "John", Age: 25}
	person.Greet()
	fmt.Println(person)
}

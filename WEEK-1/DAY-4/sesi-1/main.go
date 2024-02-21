package main

import "fmt"

type Person struct {
	Name       string
	Age        int
	Occupation string
}


func (p Person) introduce(show string) string {
	return show + fmt.Sprint(p.Name, " umur saya adalah ", p.Age, " dan jabatan saya sebagai ", p.Occupation)

}

func main() {

		// Slice of struct
allPerson := []Person{
	{"Datul", 30, "Business Owner"},
	{"Ismail", 28, "Manager"},
}
fmt.Println(allPerson)
	// Explicit Type Definition
	var student Person = Person{Name: "Alice", Age: 20, Occupation: "Software Engineer"} // Using struct literal
	fmt.Println("Student:", student)
	fmt.Println(student.introduce("Hello nama saya "))
	var student2 = Person{}
	student2.Name = "Nanda"
	student2.Age = 32
	student2.Occupation = "Collection"

	// Create instances / struct literals
	person1 := Person{Name: "John Doe", Age: 30, Occupation: "Software Engineer"}
	person2 := Person{Name: "Jane Smith", Age: 25, Occupation: "Doctor"}

	fmt.Println("Person 1:", person1)
	fmt.Println("Person 2:", person2)


	var anotherPerson *Person = &person1
	fmt.Println("Employee :", *anotherPerson)


	// Anonymous struct
	person3 := struct {
		Name       string
		Age        int
		Occupation string
	} {"Fredy", 32, "Banker"}
	fmt.Println("Person 3:", person3)

}
package main

import (
	"crypto/sha1"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"mathutil"
	"sort"
	"stringutil"
)

/*Struct identify an individual person*/
type person struct {
	personId uint64
	name     string
	age      int
}

type ByAge []person

/*Slice to store a group of people*/
type People struct {
	Group []person
}

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].age < a[j].age }

func (per *People) HirePeople(employee person) []person {
	per.Group = append(per.Group, employee)
	return per.Group
}

type MyBoxItem struct {
	Name string
}

/*Slice*/
type MyBox struct {
	Items []MyBoxItem
}

func (box *MyBox) AddItem(item MyBoxItem) []MyBoxItem {
	box.Items = append(box.Items, item)
	return box.Items
}

func main() {
	fmt.Printf("Hello, world. Now in hello.go\n")

	fmt.Printf("Calling reverse function.\n")

	fmt.Printf(stringutil.Reverse("!oG ,olleH\n"))
	fmt.Println("Calling Mulitiplication Function\n")
	fmt.Println(mathutil.Multiplication(1, 2))
	fmt.Println("Calling Division Function\n")
	fmt.Println(mathutil.Division(float32(10.0), float32(2.0)))
	/*Creating an employee*/

	employeeName := "Louie Celiberti"
	hashBytes := sha1.Sum([]byte(employeeName))
	myint := binary.BigEndian.Uint64(hashBytes[:])

	employee := person{personId: myint, name: employeeName, age: 45}

	fmt.Println(employee.name)
	fmt.Println(employee.age)
	/*Declare a ppl as a slice of persons*/
	ppl := []person{}
	/*Declare a company as a group of ppl (slice)*/
	company := People{ppl}
	company.HirePeople(employee)

	employeeName = "Raymond James"
	hashBytes = sha1.Sum([]byte(employeeName))
	myint = binary.BigEndian.Uint64(hashBytes[:])

	employee = person{personId: myint, name: employeeName, age: 12}

	/*Hiring People (adding them to the company slice)*/
	company.HirePeople(employee)
	fmt.Println("How many people are in the company:")
	fmt.Println(len(company.Group))
	fmt.Println(company.Group)
	sort.Sort(ByAge(company.Group))
	fmt.Println(company)

	/*Creating an item*/
	item1 := MyBoxItem{Name: "Test Item 1"}
	/* Declaring Items as the slice of Box*/
	items := []MyBoxItem{}
	/* Declaring a box as a box of items*/
	box := MyBox{items}

	box.AddItem(item1)
	box.AddItem(MyBoxItem{Name: "Test Item 2"})

	fmt.Println("Printing Items")
	fmt.Println(len(box.Items))

	s := "sha1 this string"
	h := sha1.New()
	h.Write([]byte(s))
	sha1_hash := hex.EncodeToString(h.Sum(nil))

	fmt.Println("Printing Hash value")
	fmt.Println(s, sha1_hash)

	fmt.Println("Printing Hash Int value")
	fmt.Println(myint)

}

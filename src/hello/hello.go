package main

import (
	"crypto/sha1"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	/*"github.com/dgryski/go-bloomindex"*/
	"datautil"
	"mathutil"
	"sort"
	"stringutil"
)

import "github.com/tylertreat/BoomFilters"

/*Struct identify an individual person*/
type person struct {
	personId uint64
	name     string
	age      int
}

type ByAge []person

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].age < a[j].age }

type persons []person

/*Slice to store a group of people*/
type People struct {
	Group []person
}

/*Populate people slice by using a "HirePeople" method */
func (per *People) HirePeople(employee person) []person {
	per.Group = append(per.Group, employee)
	return per.Group
}

func Searchstruc(company People) string {

	for _, pl := range company.Group {
		fmt.Println("ID: ", pl.personId, "Name: ", pl.name, "Age: ", pl.age)

		fmt.Println("")
	}
	return "done"

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
	newemployee := datautil.Genericperson{PersonId: myint, Name: employeeName, Age: 45}

	fmt.Println(employee.name)
	fmt.Println(employee.age)
	/*Declare a ppl as a slice of persons
	ppl := []person{}
	*/
	/*Declare a company as a group of ppl (slice)*/

	company := People{}
	company.HirePeople(employee)

	newcompany := datautil.GenericPeople{}
	newcompany.HirePeople(newemployee)

	employeeName = "Raymond James"
	hashBytes = sha1.Sum([]byte(employeeName))
	myint = binary.BigEndian.Uint64(hashBytes[:])

	employee = person{personId: myint, name: employeeName, age: 12}
	newemployee = datautil.Genericperson{PersonId: myint, Name: employeeName, Age: 12}

	/*Hiring People (adding them to the company slice)*/
	company.HirePeople(employee)
	fmt.Println("How many people are in the company:")
	fmt.Println(len(company.Group))

	fmt.Println("Printing people and ages in company from function. Pre-Sort")
	Searchstruc(company)
	sort.Sort(ByAge(company.Group))
	fmt.Println("Printing people and ages in company from function. post-Sort")
	Searchstruc(company)

	/*Hiring People in the new company using package functions*/
	newcompany.HirePeople(newemployee)
	fmt.Println("Printing people and ages of the new Company from function. Pre-Sort")
	/*Using package search function*/
	datautil.Searchstruc(newcompany)
	sort.Sort(datautil.GenericByAge(newcompany.Group))
	fmt.Println("Printing people and ages of the new Company from function. Post-Sort")
	datautil.Searchstruc(newcompany)

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

	fmt.Println("This is my bloom filter codes")

	sbf := boom.NewDefaultStableBloomFilter(10000, 0.01)
	fmt.Println("stable point", sbf.StablePoint())

	sbf.Add([]byte(`a`))

	if sbf.Test([]byte(`a`)) {
		fmt.Println("contains a")
	}

	if !sbf.TestAndAdd([]byte(`b`)) {
		fmt.Println("doesn't contain b")
	}

	if sbf.Test([]byte(`b`)) {
		fmt.Println("now it contains b!")
	}

	// Restore to initial state.
	sbf.Reset()

}

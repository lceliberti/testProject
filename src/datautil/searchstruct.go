package datautil

import "fmt"

type Genericperson struct {
	PersonId uint64
	Name     string
	Age      int
}

type GenericPeople struct {
	Group []Genericperson
}

func (per *GenericPeople) HirePeople(employee Genericperson) []Genericperson {
	per.Group = append(per.Group, employee)
	return per.Group
}

func Searchstruc(company GenericPeople) string {

	for _, pl := range company.Group {
		fmt.Println("ID: ", pl.PersonId, "Name: ", pl.Name, "Age: ", pl.Age)

		fmt.Println("")
	}
	return "done"
}

type GenericByAge []Genericperson

func (a GenericByAge) Len() int           { return len(a) }
func (a GenericByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a GenericByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

package datautil

import "fmt"
import "github.com/tylertreat/BoomFilters"

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

func InitializeBoomFilter() *boom.StableBloomFilter {
	//return boom.NewDefaultStableBloomFilter(10000, 0.01)
	return boom.NewDefaultStableBloomFilter(10000, 0.01)
}

var sbf *boom.StableBloomFilter

// Restore to initial state.
//	sbf.Reset()

func Searchstruc(company GenericPeople) string {
	sbf = InitializeBoomFilter()
	i := 0
	var dd string

	for _, pl := range company.Group {

		//	fmt.Println("ID: ", pl.PersonId, "Name: ", pl.Name, "Age: ", pl.Age)

		sbf.Add([]byte(pl.Name))
		if i == 999 {
			dd = pl.Name

		}
		//	fmt.Println("")
		i = i + 1
	}
	fmt.Println(i)
	fmt.Println(dd)
	return "done"

}

func TestBoomFilter(str string) {

	if sbf.Test([]byte(str)) {
		fmt.Println("Calling from functino, index contains: ", str)
	} else {
		fmt.Println("Calling from functino,does NOT contain: ", str)
	}
}

type GenericByAge []Genericperson

func (a GenericByAge) Len() int           { return len(a) }
func (a GenericByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a GenericByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

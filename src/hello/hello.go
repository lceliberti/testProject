package main

import (
	"bytes"
	"crypto/sha1"
	"datautil"
	"encoding/csv"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"hash/crc32"
	"io"
	"io/ioutil"
	"log"
	"mathutil"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/dgryski/go-bloomindex"
	//"reflect"
	"sort"
	//"strconv"
	"stringutil"
)

/*Struct identify an individual person*/
type person struct {
	personId uint64
	name     string
	age      int
}

type ByAge []person

type Todo1 struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	Completed bool      `json:"completed"`
	Due       time.Time `json:"due"`
}

type Todo struct {
	Id        uint64    `json:"id"`
	Name      string    `json:"name"`
	Completed bool      `json:"completed"`
	Due       time.Time `json:"due"`
}

type Todoname struct {
	Name string `json:"name"`
}

type Todosnames []Todoname

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

/*
func Searchstruc(company People) string {

	for _, pl := range company.Group {
		fmt.Println("ID:2 ", pl.personId, "Name: ", pl.name, "Age: ", pl.age)

		fmt.Println("")
	}
	return "done"

}
*/
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
	arg := os.Args[1]

	fmt.Println(arg)

	fmt.Printf("Hello, world. Now in hello.go\n")

	fmt.Printf("Calling reverse function.\n")

	fmt.Printf(stringutil.Reverse("!oG ,olleH\n"))
	fmt.Println("Calling Mulitiplication Function\n")
	fmt.Println(mathutil.Multiplication(1, 2))
	fmt.Println("Calling Division Function\n")
	fmt.Println(mathutil.Division(float32(10.0), float32(2.0)))
	/*Creating an employee*/

	employeeName := "Louie Celiberti"

	myint := stringutil.HashThisString(employeeName)

	//employee := datautil.Genericperson{personId: myint, name: employeeName, age: 45}
	employee := datautil.Genericperson{PersonId: myint, Name: employeeName, Age: 45}
	/*
		fmt.Println(employee.name)
		fmt.Println(employee.age)
	*/
	/*Declare a ppl as a slice of persons
	ppl := []person{}
	*/
	/*Declare a company as a group of ppl (slice)*/
	company := datautil.GenericPeople{}
	/*********************

			company.HirePeople(employee)
	******/
	newcompany := datautil.GenericPeople{}
	/******
		newcompany.HirePeople(newemployee)

		employeeName = "Raymond James"

		myint = stringutil.HashThisString(employeeName)

		employee = person{personId: myint, name: employeeName, age: 12}
		newemployee = datautil.Genericperson{PersonId: myint, Name: employeeName, Age: 12}
	********************/
	/*Hiring People (adding them to the company slice)*/
	/************
	company.HirePeople(employee)
	company.HirePeople(employee)
	employeeName = "Cheryl Celiberti"

	myint = stringutil.HashThisString(employeeName)

	employee = person{personId: myint, name: employeeName, age: 2}
	newemployee = datautil.Genericperson{PersonId: myint, Name: employeeName, Age: 12}

	company.HirePeople(employee)

	employeeName = "Chris James"

	myint = stringutil.HashThisString(employeeName)

	employee = person{personId: myint, name: employeeName, age: 2}
	newemployee = datautil.Genericperson{PersonId: myint, Name: employeeName, Age: 12}

	company.HirePeople(employee)

	employeeName = "Richard James"

	myint = stringutil.HashThisString(employeeName)

	employee = person{personId: myint, name: employeeName, age: 2}
	newemployee = datautil.Genericperson{PersonId: myint, Name: employeeName, Age: 12}

	company.HirePeople(employee)
	*/
	file, err := os.Open("test.csv")
	if err != nil {
		fmt.Println("Error", err)
		return
	}
	defer file.Close()

	reader := csv.NewReader(file)
	record, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error", err)
	}

	for value := range record { // for i:=0; i<len(record)
		//		fmt.Println("", record[value])
		employeeName = strings.Join(record[value], "")

		myint = stringutil.HashThisString(employeeName)

		//employee = datautil.Genericperson{personId: myint, name: employeeName, age: 2}
		employee = datautil.Genericperson{PersonId: myint, Name: employeeName, Age: 12}

		company.HirePeople(employee)

	}

	fmt.Println("How many people are in the company:")
	fmt.Println(len(company.Group))

	//	fmt.Println("Printing people and ages in company from function. Pre-Sort")
	//	Searchstruc(company)
	//sort.Sort(ByAge(company.Group))
	//	fmt.Println("Printing people and ages in company from function. post-Sort")
	//	Searchstruc(company)

	/*Hiring People in the new company using package functions*/
	//newcompany.HirePeople(newemployee)
	fmt.Println("Printing people and ages of the new Company from function. Pre-Sort")
	/*Using package search function*/
	datautil.Searchstruc(company)
	sort.Sort(datautil.GenericByAge(newcompany.Group))
	fmt.Println("Printing people and ages of the new Company from function. Post-Sort")
	//	datautil.Searchstruc(newcompany)

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
	/*
		sbf := boom.NewDefaultStableBloomFilter(10000, 0.01)
		fmt.Println("stable point", sbf.StablePoint())

		sbf.Add([]byte(employeeName))
	*/

	datautil.TestBoomFilter(arg)
	datautil.TestBoomFilter(employeeName)

	idx := bloomindex.NewShardedIndex(.000000001, 10)
	//idx := bloomindex.NewShardedIndex(.01, 4)

	/*populate BloomIndex, by looking through company*/
	//	var toks []uint32
	//	var qtoks []uint32
	//var qtoks2 []uint32
	var m map[uint64]uint64
	m = make(map[uint64]uint64)
	var iter uint64
	/***************************
		iter = 0
		for _, pl := range company.Group {
			m[iter] = pl.personId
			fmt.Println("ID: ", pl.personId, "Name: ", pl.name, "Age: ", pl.age)
			tokens := strings.Fields(pl.name)
			var toks []uint32

			for _, t := range tokens {
				fmt.Println("Name: ", pl.name, "token: ", t)
				toks = append(toks, crc32.ChecksumIEEE([]byte(t)))
			}
			idx.AddDocument(toks)
			fmt.Println("Docid:", pl.name)
			fmt.Println(crc32.ChecksumIEEE([]byte(pl.name)))
			iter = iter + 1
		}
	*********/

	iter = 0
	for _, pl := range company.Group {
		m[iter] = pl.PersonId
		//fmt.Println("ID: ", pl.personId, "Name: ", pl.name, "Age: ", pl.age)
		tokens := strings.Fields(pl.Name)
		var toks []uint32

		for _, t := range tokens {
			//fmt.Println("Name: ", pl.name, "token: ", t)
			t = strings.Replace(t, ",", "", 5)
			toks = append(toks, crc32.ChecksumIEEE([]byte(t)))
			for ngram, frequency := range Parse(t, 10) {
				frequency = frequency + frequency
				//				fmt.Println("Name: ", pl.name, "token: ", ngram)
				ngram = strings.Replace(ngram, " ", "", 5)
				toks = append(toks, crc32.ChecksumIEEE([]byte(ngram)))
			}

		}
		idx.AddDocument(toks)
		//fmt.Println("Docid:", pl.name)
		//fmt.Println(crc32.ChecksumIEEE([]byte(pl.name)))
		iter = iter + 1
	}

	var toks []uint32
	qstr := arg
	//toks = append(toks, crc32.ChecksumIEEE([]byte("Fund")))
	toks = append(toks, crc32.ChecksumIEEE([]byte(qstr)))

	//toks = append(toks, crc32.ChecksumIEEE([]byte("Premi")))

	/*
		query := []string{qstr}


		for _, q := range query {
			toks = append(toks, crc32.ChecksumIEEE([]byte(q)))
		}
	*/
	ids := idx.Query(toks)

	fmt.Println("Printing ids returned from", qstr)
	fmt.Println("length of ids: ", len(ids))
	fmt.Println(ids)
	fmt.Println(ids[0])

	var myToDoID uint64
	var count int
	count = 0
	for _, doc := range ids {
		/*
			fmt.Println("printing docs for", qstr)
			fmt.Println(strconv.FormatUint(uint64(doc), 16))
			fmt.Println(m[uint64(doc)])
		*/
		myToDoID = m[uint64(doc)]

		//	fmt.Println("Calling getToDos for: ", myToDoID)
		getToDos(myToDoID)
		count = count + 1
	}
	fmt.Println("Count of tods : ", count)
	/******* Calling findToDo*****/
}

func Parse(text string, n int) map[string]int {
	chars := []rune(strings.Repeat(" ", n))
	table := make(map[string]int)

	for _, letter := range strings.Join(strings.Fields(text), " ") + " " {
		chars = append(chars[1:], letter)

		ngram := string(chars)
		if _, ok := table[ngram]; ok {
			table[ngram]++
		} else {
			table[ngram] = 1
		}
	}

	return table
}

func getToDos(myToDoID uint64) {
	url := "http://localhost:8080/findtodos"
	//	fmt.Println("URL:>", url)
	tdn := Todo{Id: myToDoID}
	//	fmt.Println(myToDoID)

	var tdns []Todo
	//populate slice by appending with struct
	tdns = append(tdns, tdn)
	jsonStr, err := json.Marshal(tdns)

	//Post Request
	//	fmt.Println(jsonStr)
	req, err := http.NewRequest("GET", url, bytes.NewBuffer(jsonStr))

	req.Header.Set("X-Custom-Header", "myvalue")
	req.Header.Set("Content-Type", "application/json")

	//Create an HTTP client and execute request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	//defer resp.Body.Close()

	//	fmt.Println("response Status:", resp.Status)
	//	fmt.Println("response Headers:", resp.Header)

	body, err := ioutil.ReadAll(io.LimitReader(resp.Body, 1048576))
	if err != nil {
		panic(err)
	}

	//	fmt.Println(os.Stdout, string(body)) //<-- here !

	var todos []Todo

	if err = json.Unmarshal(body, &todos); err != nil {
		log.Println("Error unmarshelling")
		log.Println(err)
	}

	for _, tasks := range todos {
		//	fmt.Println("ID: ", tasks.Id, "Name: ", tasks.Name, "Complete: ", tasks.Completed)
		fmt.Println(tasks.Name)

	}

}

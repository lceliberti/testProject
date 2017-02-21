package main

import (
	"fmt"
	"hash/crc32"
	"log"
	"strings"
	"stringutil"

	bloomindex "github.com/dgryski/go-bloomindex"
)

var idx *bloomindex.ShardedIndex

var m map[uint64]uint64

var currentId int

var todos Todos

// Give us some seed data
func init() {
	m = make(map[uint64]uint64)
	log.Println("Creating Index")
	CreateIndex()
	RepoCreateTodo(Todo{Name: "Write 2nd presentation"})
	RepoCreateTodo(Todo{Name: "Schedule 2nd meeting"})

}

func RepoFindTodo(id uint64) Todos {
	log.Println("calling repofindtodo with: ", id)
	var retTodos Todos
	for _, t := range todos {
		if t.Id == id {
			log.Println("Found Name ", t.Name)
			retTodos = append(retTodos, t)
		}
	}
	// return empty Todo if not found
	return retTodos
}

func RepoCreateTodo(t Todo) Todo {
	//currentId += 1
	var currentId uint64
	var toDoID uint64

	currentId = stringutil.HashThisString(string(t.Name))
	t.Id = currentId
	fmt.Println("Creating Task: ")
	fmt.Println(t.Name)
	t.Name = t.Name //+ stringutil.Reverse("!oG ,olleH\n")
	todos = append(todos, t)
	toDoID = RepoPopulateIndex(string(t.Name))
	m[toDoID] = currentId

	for k, v := range m {
		fmt.Println("k:", k, "v:", v)
	}

	return t
}

func RepoDestroyTodo(id uint64) error {
	for i, t := range todos {
		if t.Id == id {
			todos = append(todos[:i], todos[i+1:]...)
			return nil
		}
	}
	log.Println("Could not find Todo with id of %d to delete", id)
	return fmt.Errorf("Could not find Todo with id of %d to delete", id)

}

func CreateIndex() {
	idx = bloomindex.NewShardedIndex(0.01, 4)
}

func RepoPopulateIndex(name string) uint64 {

	var myToDoID uint64

	//m[iter] = pl.personId
	tokens := strings.Fields(string(name))
	var toks []uint32
	var qtoks []uint32
	var hName uint32
	var checkToks []uint32
	checkToks = append(qtoks, crc32.ChecksumIEEE([]byte(name)))

	checkids := idx.Query(checkToks)

	if len(checkids) > 0 {
		fmt.Println("Document Exists, returning existing DOCID")
		for _, doc := range checkids {

			fmt.Println(doc)
			myToDoID = uint64(doc)
		}
		return myToDoID
	}

	for _, t := range tokens {
		fmt.Println("Name: ", name, "Token: ", t)
		toks = append(toks, crc32.ChecksumIEEE([]byte(t)))
	}
	hName = crc32.ChecksumIEEE([]byte(name))
	toks = append(toks, hName)
	log.Println("Populating index for: ", name)

	idx.AddDocument(toks)

	fmt.Println("Docid:", name)

	fmt.Println(crc32.ChecksumIEEE([]byte(name)))

	qtoks = append(qtoks, crc32.ChecksumIEEE([]byte(name)))

	ids := idx.Query(qtoks)

	fmt.Println(len(qtoks))
	fmt.Println(len(ids))

	for _, doc := range ids {

		fmt.Println(doc)
		myToDoID = uint64(doc)
	}

	return myToDoID

}

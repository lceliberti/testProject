package main

import (
	"fmt"

	"github.com/rtt/Go-Solr"
)

/*
 * README
 * ------
 * This example shows a Query being performed. A query is built up using
 * a Query type object, the query executed and the results are then
 * printed to the console
 */

func main() {

	// init a connection

	s, err := solr.Init("localhost", 8983, "techproducts")

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Running Query")
	// Build a query object
	// Here we are specifying a 'q' param,
	// rows, faceting and facet.fields
	///http://localhost:8983/solr/techproducts/update/extract?literal.id=doc7&literal.Manu=louie&commit=true -F "myfile=@example/exampledocs/solr-word.pdf"
	q := solr.Query{
		Params: solr.URLParamMap{
			"q": []string{"content:fou"},
			/*"facet.field": []string{"accepts_4x4s", "accepts_bicycles"},
			"facet":       []string{"true"},*/
		},
		Rows: 10,
	}

	// perform the query, checking for errors
	res, err := s.Select(&q)

	if err != nil {
		fmt.Println(err)
		return
	}

	// grab results for ease of use later on
	results := res.Results

	// print a summary and loop over results, priting the "title" and "latlng" fields
	fmt.Println(
		fmt.Sprintf("Query: %#v\nHits: %d\nNum Results: %d\nQtime: %d\nStatus: %d\n\nResults\n-------\n",
			q,
			results.NumFound,
			results.Len(),
			res.QTime,
			res.Status))

	for i := 0; i < results.Len(); i++ {
		fmt.Println("ID:", results.Get(i).Field("id"))
		fmt.Println("Name:", results.Get(i).Field("name"))
		fmt.Println("Manu:", results.Get(i).Field("manu"))
		fmt.Println("Features:", results.Get(i).Field("features"))
		fmt.Println("Content:", results.Get(i).Field("content"))
		fmt.Println("")
	}

}

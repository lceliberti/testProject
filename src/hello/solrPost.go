package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/rtt/Go-Solr"
)

type doc struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	manu     string `json:"manu"`
	features string `json:"featurs"`
}

type docs []doc

// -----------

type Name struct {
	Set string `json:"set"`
}

type updatedoc struct {
	Id   string `json:"id"`
	Name Name   `json:"name"`
}

type updatedocs []updatedoc

/*
 * README
 * ------
 * This example shows a Query being performed. A query is built up using
 * a Query type object, the query executed and the results are then
 * printed to the console
 */

func main() {
	url := "http://localhost:8983/solr/techproducts/update/json?commit=true"
	//http://localhost:8983/solr/techproducts/update/extract?stream.file=./dealfolders/test23.docx&literal.id=doc52&literal.name=louieceliberti25&literal.manu=mymanufacturer&commit=true

	setDocInSolr(url)
	updateAttributeInSolr(url)
	querySOLR()
}

func setDocInSolr(url string) {
	vdoc := doc{Id: "doc55", Name: "louiecelibertiGOUpdate", manu: "bbbb"}
	//Declare a variable of slice type
	var vdocs docs
	//populate slice by appending with struct
	vdocs = append(vdocs, vdoc)

	jsonStr, err := json.Marshal(vdocs)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))

	req.Header.Set("X-Custom-Header", "myvalue")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	//defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))

}

func updateAttributeInSolr(url string) {
	/*Update document with new name*/
	var upd updatedoc
	upd = updatedoc{Id: "doclwc",
		Name: Name{Set: "louiecelibertiupdatedGO Gators!!!"},
	}

	var vupdatedocs updatedocs
	//populate slice by appending with struct
	vupdatedocs = append(vupdatedocs, upd)

	jsonStr, _ := json.Marshal(vupdatedocs)

	fmt.Println("Printing jsonstr: ", upd)

	req2, _ := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))

	req2.Header.Set("Content-Type", "application/json")

	client2 := &http.Client{}
	fmt.Println("Calling DO")
	/*
		reqbody, _ := ioutil.ReadAll(req2.Body)
		fmt.Println("response Body:", string(reqbody))
	*/
	resp2, er2 := client2.Do(req2)
	if er2 != nil {
		panic(er2)
	}

	fmt.Println("Called DO")
	fmt.Println("Printing Body of response: ", resp2.Body)

	fmt.Println("response Status:", resp2.Status)
	fmt.Println("response Headers:", resp2.Header)
	body2, _ := ioutil.ReadAll(resp2.Body)
	fmt.Println("response Body:", string(body2))
}

func querySOLR() {
	//defer resp.Body.Close()
	fmt.Println("Querying SOLR")

	// init a connection

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

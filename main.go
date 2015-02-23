// neoism2 project main.go
package main

import (
	"fmt"
	"github.com/jmcvetta/neoism"
)

type Page struct {
	Title string
	Body  string
}

func (p *Page) save() error {
	db, err := neoism.Connect("http://localhost:7474/db/data")
	if err != nil {
		return err
	}
	res := []struct {
		N neoism.Node
	}{}
	cq := neoism.CypherQuery{
		Statement:  "MERGE (n:Page {title: {title}}) ON MATCH SET n.body = {body} RETURN n",
		Parameters: neoism.Props{"title": p.Title, "body": p.Body},
		Result:     &res,
	}
	db.Cypher(&cq)
	return nil
}

func load(title string) (*Page, error) {
	db, err := neoism.Connect("http://localhost:7474/db/data")
	if err != nil {
		return nil, err
	}
	cr := []struct {
		Title string `json:"n.title"`
		Body  string `json:"n.body"`
	}{}
	cq := neoism.CypherQuery{
		Statement:  "MATCH (n:Page) WHERE n.title = {title} RETURN n.title, n.body",
		Parameters: neoism.Props{"title": title},
		Result:     &cr,
	}
	db.Cypher(&cq)
	r := cr[0]
	fmt.Println(r.Body, r.Title)
	return &Page{Title: r.Title, Body: r.Body}, nil
}

func main() {
	q := Page{Title: "asian", Body: "nothing"}
	q.save()
	r, _ := load("asian")
	fmt.Println(r.Title, r.Body)
}

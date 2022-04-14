package main

import (
	"bytes"
	"encoding/json"
	es "github.com/elastic/go-elasticsearch/v7"
	"github.com/stretchr/testify/assert"
	"log"
	"net/http"
	"strings"
	"testing"
)

var client *es.Client

func init() {
	var err error
	client, err = es.NewClient(es.Config{
		Addresses: []string{"http://dev.ll:9200"},
		Username:  "elastic",
		Password:  "123456",
	})

	if err != nil {
		log.Fatalln(err)
	}
}

func TestCheckIndexExists(t *testing.T) {
	rsp, err := client.Indices.Get([]string{"acger"})

	if err != nil {
		t.Fatal(err)
	}

	if rsp.StatusCode == http.StatusNotFound {
		t.Log("not found")
	} else {
		t.Log("found")
	}
}

type Made struct {
	Name   string  `json:"name,omitempty"`
	Weight float32 `json:"weight,omitempty"`
}

type doc struct {
	Doc interface{} `json:"doc"`
}

func TestUpdateDocument(t *testing.T) {
	a := assert.New(t)

	body := &bytes.Buffer{}

	err := json.NewEncoder(body).Encode(&doc{
		Doc: &Made{
			Weight: 31.11,
		},
	})

	a.Nil(err)

	rsp, err := client.Update("acger", "001", body)

	a.Nil(err)

	t.Log(rsp)
}

func TestSearch(t *testing.T){
	a := assert.New(t)

	body := &bytes.Buffer{}

	body.WriteString(`
	  {
		"query": {
		"match_all": {}
		}
	  }
	`)

	rsp, err := client.Search(client.Search.WithIndex("acger"), client.Search.WithBody(body))

	a.Nil(err)

	t.Log(rsp)
}

func TestUpsertDocument(t *testing.T) {
	a := assert.New(t)
	body := &bytes.Buffer{}

	err := json.NewEncoder(body).Encode(&Made{
		Name:   "mimi",
		Weight: 32.2,
	})

	a.Nil(err)

	t.Log("upsert")
	rsp, err := client.Index("acger", body, client.Index.WithDocumentID("004"))
	a.Nil(err)
	t.Log(rsp)

	t.Log("get")
	rsp, err = client.Get("acger", "004")
	a.Nil(err)
	t.Log(rsp)

	t.Log("delete")
	rsp, err = client.Delete("acger", "004")
	a.Nil(err)
	t.Log(rsp)
}

func TestCreateDocument(t *testing.T) {
	a := assert.New(t)
	body := &bytes.Buffer{}

	err := json.NewEncoder(body).Encode(&Made{
		Name:   "illya",
		Weight: 30.2,
	})

	a.Nil(err)

	rsp, err := client.Create("acger", "001", body)
	a.Nil(err)

	t.Log(rsp)
}

func TestCreateSimpleIndex(t *testing.T) {
	rsp, err := client.Indices.Create("acger", client.Indices.Create.WithBody(strings.NewReader(`
	{
	  "settings": {
		"number_of_replicas": 0,
		"number_of_shards": 5
	  },
	  "mappings": {
		"properties": {
		  "name" : {
			"type": "keyword"
		  },
		  "weight":{
			"type" : "half_float"
		  }
		}
	  }
	}
	`)))

	if err != nil {
		t.Fatal(err)
	}

	t.Log(rsp)
}

func TestCreateIndex(t *testing.T) {
	a := assert.New(t)
	response, err := client.Indices.Create("book-0.1.0", client.Indices.Create.WithBody(strings.NewReader(`
	{
		"aliases": {
			"book":{}
		},
		"settings": {
			"analysis": {
				"normalizer": {
					"lowercase": {
						"type": "custom",
						"char_filter": [],
						"filter": ["lowercase"]
					}
				}
			}
		},
		"mappings": {
			"properties": {
				"name": {
					"type": "keyword",
					"normalizer": "lowercase"
				},
				"price": {
					"type": "double"
				},
				"summary": {
					"type": "text",
					"analyzer": "ik_max_word"
				},
				"author": {
					"type": "keyword"
				},
				"pubDate": {
					"type": "date"
				},
				"pages": {
					"type": "integer"
				}
			}
		}
	}
	`)))
	a.Nil(err)
	t.Log(response)
}



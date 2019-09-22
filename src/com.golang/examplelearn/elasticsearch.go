package main

import "fmt"

/**
elasticsearch demo
*/
import (
	"gopkg.in/olivere/elastic.v2"
)

type Tweet struct {
	User    string
	Message string
}

func main() {
	client, err := elastic.NewClient(elastic.SetSniff(false), elastic.SetURL("http://hadoop-senior-1.cathome.com:9200/"))
	if err != nil {
		fmt.Println("connect es error", err)
		return
	}

	fmt.Println("connect elasticserch success")

	tweet := Tweet{User: "olivere", Message: "Take Five"}
	_, err = client.Index().
		Index("twitter").
		Type("tweet").
		Id("1").
		BodyJson(tweet).
		Do()
	if err != nil {
		panic(err)
		return
	}

	fmt.Println("insert success")

}

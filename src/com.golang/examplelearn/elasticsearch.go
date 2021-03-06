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

	for i := 1; i <= 10; i++ {
		_, err = client.Index().
			Index("twitter").
			Type("tweet").
			Id(fmt.Sprintf("%d", i)).
			BodyJson(tweet).
			Do()
		if err != nil {
			panic(err)
			return
			fmt.Sprintf("insert %d success", i)
		}
	}

}

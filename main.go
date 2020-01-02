package main 

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func swap(a, b *int) {
	*a = *a + *b
	*b = *a - *b
	*a = *a - *b
}

func main() {
	if len(os.Args) < 5 {
		log.Fatal("usage: art [ip address] [repo] [username] [password]")
	}
	log.Println("Art v0.0.1")
	log.Println(os.Args)

	client := NewClient(os.Args[3], os.Args[4])

	artifacts := client.GetArtifactList()

	first := 0
	second := 0

	largest := 0
	secondLargest := 0

	for i := 0; i < len(artifacts.Results); i++ {
		stat := GetFileStats(artifacts.Results[i].PathName())
		if stat.DownloadCount > second {
			second = stat.DownloadCount
		}
		if second > first {
			first = first + second
			second = first - second
			first = first - second

			largest = largest + secondLargest
			secondLargest = largest - secondLargest
			largest = largest - secondLargest
		}
	}

	fmt.Printf("Second largest file: %s with %d downloads\n", artifacts.Results[secondLargest].PathName(), second)
}

func GetFileStats(uri string) *ArtifactoryFileStats {
	url := fmt.Sprintf("http://%s/artifactory/api/storage/%s/%s?stats", os.Args[1], os.Args[2], uri)
		
	resp, err := http.Get(url)
	if nil != err {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	bytes, err := ioutil.ReadAll(resp.Body)
	if nil != err {
		log.Fatal(err)
	}

	stats := &ArtifactoryFileStats{}

	err = json.Unmarshal(bytes, stats)
	if nil != err {
		log.Fatal(err)
	}

	return stats
}

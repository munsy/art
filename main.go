package main 

import (
	"fmt"
	"log"
	"os"
	"sort"
)

func main() {
	if len(os.Args) < 5 {
		log.Fatal("usage: art [ip address] [repo] [username] [password]")
	}
	log.Println("Art v0.0.1")
	log.Println(os.Args)

	client := NewClient(os.Args[3], os.Args[4])
	
	artifacts, err := client.GetArtifactList()

	if nil != err {
		log.Fatal(err)
	}

	sort.Stable(artifacts)

	for i := 0; i < len(artifacts.Results); i++ {
		fmt.Printf("%d\t%s\n", artifacts.Results[i].TotalDownloads(), artifacts.Results[i].Name)
	}
}


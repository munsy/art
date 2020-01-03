package main 

import (
	"fmt"
	"log"
	"os"
	"sort"
)

const (
	major = 1
	minor = 0
	patch = 0
)

func main() {
	log.Printf("Art v%d.%d.%d\n", major, minor, patch)
	if len(os.Args) < 5 {
		log.Fatal("usage: art [ip address] [repo] [username] [password]")
	}

	client, err := NewClient(os.Args[3], os.Args[4])
	
	if nil != err {
		log.Fatal(err)
	}

	artifacts, err := client.GetArtifactList(os.Args[2])

	if nil != err {
		log.Fatal(err)
	}

	sort.Stable(artifacts)

	for i := 0; i < len(artifacts.Results); i++ {
		fmt.Printf("%d\t%s\n", artifacts.Results[i].TotalDownloads(), artifacts.Results[i].Name)
	}
}


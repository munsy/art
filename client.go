package main

import(
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

type ArtClient struct {
	client *http.Client
	basicAuth string
}

func NewClient(username, password string) *ArtClient {
	c := &http.Client{
		Timeout: 5 * time.Second,
	}

	auth := fmt.Sprintf("%s:%s", os.Args[3], os.Args[4])
	base := base64.StdEncoding.EncodeToString([]byte(auth))
	header := fmt.Sprintf("Basic %s", base)

	return &ArtClient{
		client: c,
		basicAuth: header,
	}
}

func (a *ArtClient) GetArtifactList() *ArtifactList {
	url := fmt.Sprintf("http://%s/artifactory/api/search/aql", os.Args[1])
	aql := fmt.Sprintf(`items.find({"repo":{"$eq":"%s"}})`, os.Args[2])
	reader := strings.NewReader(aql)
	
	post, err := http.NewRequest("POST", url, reader)
	if nil != err {
		log.Fatal(err)
	}

	post.Header.Add("Authorization", a.basicAuth)

	resp, err := a.client.Do(post)
	if nil != err {
		log.Fatal(err)
	}

	if nil != err {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	out, err := ioutil.ReadAll(resp.Body)
	if nil != err {
		log.Fatal(err)
	}

	log.Println("Received valid response. Decoding list...")

	artifacts := &ArtifactList{}

	err = json.Unmarshal(out, artifacts)
	if nil != err {
		log.Fatal(err)
	}

	if nil == artifacts {
		log.Fatal("nil artifact list")
	}

	return artifacts
}
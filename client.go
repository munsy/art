package main

import(
	"encoding/base64"
	"encoding/json"
	"errors"
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

func NewClient(username, password string) (*ArtClient, error) {
	if len(username) == 0 {
		return nil, errors.New("no username supplied")
	}
	if len(password) == 0 {
		return nil, errors.New("no password supplied")
	}

	c := &http.Client{
		Timeout: 5 * time.Second,
	}

	auth := fmt.Sprintf("%s:%s", username, password)
	base := base64.StdEncoding.EncodeToString([]byte(auth))
	header := fmt.Sprintf("Basic %s", base)

	return &ArtClient{
		client: c,
		basicAuth: header,
	}, nil
}

func (a *ArtClient) GetArtifactList(repo string) (*ArtifactList, error){
	url := fmt.Sprintf("http://%s/artifactory/api/search/aql", os.Args[1])
	aql := fmt.Sprintf(`items.find({"repo":{"$eq":"%s"}}).include("stat")`, repo)
	reader := strings.NewReader(aql)
	
	post, err := http.NewRequest("POST", url, reader)
	if nil != err {
		return nil, err
	}

	post.Header.Add("Authorization", a.basicAuth)

	resp, err := a.client.Do(post)
	if nil != err {
		return nil, err
	}

	if nil != err {
		return nil, err
	}
	defer resp.Body.Close()

	out, err := ioutil.ReadAll(resp.Body)
	if nil != err {
		return nil, err
	}

	log.Println("Received valid response. Decoding list...")

	artifacts := &ArtifactList{}

	err = json.Unmarshal(out, artifacts)
	if nil != err {
		return nil, err
	}

	if nil == artifacts {
		return nil, errors.New("nil artifact list")
	}

	return artifacts, nil
}

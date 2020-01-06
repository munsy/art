package client

import(
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

type ArtClient struct {
	client *http.Client
	basicAuth string
	URL string
}

func NewClient(url string) *ArtClient {
	c := &http.Client{
		Timeout: 5 * time.Second,
	}

	return &ArtClient{
		client: c,
		basicAuth: "",
		URL: url,
	}
}

func (a *ArtClient) SetAuth(username, password string) error {
	if len(username) == 0 {
		return errors.New("no username supplied")
	}
	if len(password) == 0 {
		return errors.New("no password supplied")
	}
	auth := fmt.Sprintf("%s:%s", username, password)
	base := base64.StdEncoding.EncodeToString([]byte(auth))
	header := fmt.Sprintf("Basic %s", base)
	a.basicAuth = header
	return nil
}

func (a *ArtClient) GetArtifactList(repo string) (*ArtifactList, error){
	url := fmt.Sprintf("http://%s/artifactory/api/search/aql", a.URL)
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

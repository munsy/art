package main 

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	if len(os.Args) < 5 {
		log.Fatal("usage: art [ip address] [repo] [username] [password]")
	}
	log.Println("Art v0.0.1")
	log.Println(os.Args)
	
	url := fmt.Sprintf("http://%s/artifactory/api/search/aql", os.Args[1])
	aql := fmt.Sprintf(`items.find({"repo":{"$eq":"%s"}})`, os.Args[2])
	auth := fmt.Sprintf("%s:%s", os.Args[3], os.Args[4])
	base := base64.StdEncoding.EncodeToString([]byte(auth))
	header := fmt.Sprintf("Basic %s", base)
	reader := strings.NewReader(aql)
	
	//resp, err := http.Post(url, "text/html", reader)
	client := &http.Client{
		Timeout: 10 * time.Second,
	}
	
	post, err := http.NewRequest("POST", url, reader)
	if nil != err {
		log.Fatal(err)
	}

	post.Header.Add("Authorization", header)

	resp, err := client.Do(post)
	if nil != err {
		log.Fatal(err)
	}

	if nil != err {
		log.Fatal(err)
	}
	defer resp.Body.Close()

//	var v interface{}
//	err = json.Unmarshal(resp.Body, &v)

	out, err := ioutil.ReadAll(resp.Body)
	if nil != err {
		log.Fatal(err)
	}

	log.Println(string(out))
	log.Println("Done!")
}

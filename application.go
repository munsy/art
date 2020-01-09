package main 

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"sort"
	"strings"

	"github.com/urfave/cli/v2"

	"github.com/munsy/art/client"
)

const (
	program = "Art"
	major = 1
	minor = 0
	patch = 0
)

func main() {
	app := &cli.App{
    	Name: "art",
    	Usage: "A quick, simple tool to grab certain data from an artifactory instance.",
    	EnableBashCompletion: true,
    	Action: func(c *cli.Context) error {
    		log.Printf("%s v%d.%d.%d\n", program, major, minor, patch)
    		log.Printf("No arguments supplied.\n")
    		return serve(c)
    	},
    	Commands: []*cli.Command{
			{
				Name:    "serve",
				Aliases: []string{"s"},
				Usage:   "Run REST API",
				Flags: 	 []cli.Flag{
		    		&cli.IntFlag{
		    			Name:    "port",
		    			Value:   5000,
		    			Aliases: []string{"p"},
		    			Usage:   "REST API port number",
		    		},
		    		&cli.BoolFlag{
		    			Name:    "container",
		    			Value:   true,
		    			Aliases: []string{"c"},
		    			Usage:   "Run in container mode",
		    		},
		    		&cli.BoolFlag{
		    			Name:    "angular",
		    			Value:   true,
		    			Aliases: []string{"a", "ng", "ui"},
		    			Usage:   "Enable Angular frontend",
		    		},
		    	},
			  	Action:  func(c *cli.Context) error {
			    	return serve(c)
			  	},
			},
			{
			  	Name:    "lookup",
			  	Aliases: []string{"l"},
			  	Usage:   "Search via command line",
			  	Flags: 	 []cli.Flag{
		    		&cli.StringFlag{
		    			Name:    "host",
		    			Value:   "art.munsy.io",
		    			Aliases: []string{"H"},
		    			Usage:   "Target host",
		    		},
		    		&cli.StringFlag{
		    			Name:    "repo",
		    			Value:   "jcenter-cache",
		    			Aliases: []string{"r"},
		    			Usage:   "Target repository",
		    		},
		    		&cli.StringFlag{
		    			Name:    "username",
		    			Value:   "",
		    			Aliases: []string{"u"},
		    			Usage:   "Username for basic authorization",
		    		},
		    		&cli.StringFlag{
		    			Name:    "password",
		    			Value:   "",
		    			Aliases: []string{"p"},
		    			Usage:   "Password for basic authorization",
		    		},
		    	},
			  	Action:  func(c *cli.Context) error {
			  	  	return lookup(c)
			  	},
			},
    	},
    }

	err := app.Run(os.Args)

	if err != nil {
		log.Fatal(err)
	}
}

type quickapi struct{}
type Quickreq struct{
	Repo 	 string `json:"repo"`
	Url 	 string `json:"url"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func (q *quickapi) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Methods", "*")
    w.Header().Set("Access-Control-Allow-Credentials", "true")
    w.Header().Set("Access-Control-Allow-Headers", "Accept, Accept-Encoding, Authorization, Content-Length, Content-Type, X-CSRF-Token")
	w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)

	switch r.Method {
	case http.MethodPost:
		if err := r.ParseForm(); nil != err {
			http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }

        var q Quickreq

        err := json.NewDecoder(r.Body).Decode(&q)
        if nil != err {
        	http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        
        c := client.NewClient(q.Url)
        if err := c.SetAuth(q.Username, q.Password); nil != err {
        	http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }

        artifacts, err := c.GetArtifactList(q.Repo)
		if nil != err {
			log.Fatal(err)
		}

		sort.Stable(artifacts)

		json.NewEncoder(w).Encode(artifacts)
		break
	default:
		http.Error(w, "404 not found", http.StatusNotFound)
		break
	}
}

func serve(c *cli.Context) error {
	http.Handle("/api/v1/artifactory", &quickapi{})
	
	if c.Bool("angular") {
		//runPrint("npm", "i")
		if c.Bool("container") {
			runPrint("./node_modules/.bin/ng", "build", "--configuration=docker")
		} else {
			runPrint("ng", "build")
		}
	}
	http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("./dist/site"))))

	port := c.Int("port")
	if port == 0 {
		port = 5000
	}

	addr := "0.0.0.0"
	binding := fmt.Sprintf("%s:%d", addr, port)
	log.Printf("API server now listening on port %s (press Ctrl^C quit)", binding)
    log.Fatal(http.ListenAndServe(binding, nil))
    return nil
}

func lookup(c *cli.Context) error {
	log.Printf("Querying repository %s on host %s\n", c.String("repo"), c.String("host"))
	look := client.NewClient(c.String("host"))
	err := look.SetAuth(c.String("username"), c.String("password"))
	
	if nil != err {
		return err
	}

	artifacts, err := look.GetArtifactList(c.String("repo"))

	if nil != err {
		return err
	}

	sort.Stable(artifacts)

	// largest download counts
	first := 0
	second := 0

	biggest := 0 // longest name length

	for i := 0; i < len(artifacts.Results); i++ {
		if artifacts.Results[i].TotalDownloads() > first {
			second = first
			first = artifacts.Results[i].TotalDownloads()
		}
		length := len(artifacts.Results[i].Name)
		if length > biggest {
			biggest = length
		}
	}

	var pop1 = make([]client.ArtifactResult, 0)
	var pop2 = make([]client.ArtifactResult, 0)

	if len(artifacts.Results) > 2 && first != second {
		fmt.Printf("There were multiple files with the same download count. Displaying files by first and second highest download counts:\n")
	}

	for i := 0; i < len(artifacts.Results); i++ {
		if artifacts.Results[i].TotalDownloads() == first {
			pop1 = append(pop1, artifacts.Results[i])
			continue
		}
		if artifacts.Results[i].TotalDownloads() == second {
			pop2 = append(pop1, artifacts.Results[i])
		}
	}

	for i := 0; i < len(pop1); i++ { 
		line := fmt.Sprintf("%d - %s", pop1[i].TotalDownloads(), pop1[i].Name)
		if i < len(pop2) {
			line = fmt.Sprintf("%-*s%d - %s", biggest, line, pop2[i].TotalDownloads(), pop2[i].Name)
		}
		fmt.Println(line)
	}

	return nil
}

func clean() {
	remove()
}

func remove(paths ...string) {
	for _, path := range paths {
		log.Println("rm -r", path)
		os.RemoveAll(path)
	}
}

func move(source, destination string) {
	log.Println("mv", source, destination)

	err := os.Rename(source, destination)
	if err != nil {
		log.Fatal(err)
	}
}

func runPrint(cmd string, args ...string) {
	log.Println(cmd, strings.Join(args, " "))
	ecmd := exec.Command(cmd, args...)
	ecmd.Stdout = os.Stdout
	ecmd.Stderr = os.Stderr
	err := ecmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}

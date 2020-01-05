package main 

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"sort"

	"github.com/urfave/cli/v2"

	"github.com/munsy/art/client"
)

const (
	program = "Art"
	major = 1
	minor = 0
	patch = 0
)

type quickapi struct{}

func (q *quickapi) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		if err := r.ParseForm(); nil != err {
			http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }

        repo := r.FormValue("repo")
        url := r.FormValue("url")
        username := r.FormValue("username")
        password := r.FormValue("password")
        
        c := client.NewClient(url)
        if err := c.SetAuth(username, password); nil != err {
        	http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }

        artifacts, err := c.GetArtifactList(repo)
		if nil != err {
			log.Fatal(err)
		}

		sort.Stable(artifacts)
    	
    	w.WriteHeader(http.StatusOK)
    	w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(artifacts)
		break
	default:
		http.Error(w, "404 not found", http.StatusNotFound)
		break
	}
}

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

func serve(c *cli.Context) error {
	http.Handle("/api/v1/artifactory", &quickapi{})
	log.Printf("API server now listening on port %d (press Control^C to stop)", c.Int("port"))
    log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", c.Int("port")), nil))
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

	for i := 0; i < len(artifacts.Results); i++ {
		fmt.Printf("%d\t%s\n", artifacts.Results[i].TotalDownloads(), artifacts.Results[i].Name)
	}
	return nil
}

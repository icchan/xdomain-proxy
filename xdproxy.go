package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
)

func main() {
	// command line args
	listenHost := flag.String("listen", ":1337", "host and port to listen on")
	rootPath := flag.String("root", "/", "relative path for apis")
	apiPath := flag.String("api", "/api", "relative path for apis")
	apiServer := flag.String("server", "", "remote path for apis")
	flag.Parse()

	// Host files in the current folder
	pwd, err := os.Getwd()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	fmt.Println(pwd)

	fsWeb := http.FileServer(http.Dir(pwd))
	http.Handle(*rootPath, http.StripPrefix(*rootPath, fsWeb))

	// proxy the api requests
	backendUrl, _ := url.Parse(*apiServer)
	http.Handle(*apiPath, httputil.NewSingleHostReverseProxy(backendUrl))

	// fire up the server
	log.Printf("Hosting current path as %s on %s ", *rootPath, *listenHost)
	log.Printf("  and forwarding %s to %s", *apiPath, *apiServer)
	http.ListenAndServe(*listenHost, nil)
}

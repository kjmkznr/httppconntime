package main

import (
	"flag"
	"github.com/kjmkznr/httppconntime"
	"log"
	"net/url"
	"os"
	"time"
)

var logger *log.Logger

func init() {
	logger = log.New(os.Stdout, "", log.LstdFlags)
}

func main() {
	var targetURL string
	var initialWaitTime time.Duration
	var maxWaitTime time.Duration

	flag.StringVar(&targetURL, "url", "", "Target URL. like http://www.example.jp/")
	flag.DurationVar(&initialWaitTime, "init", time.Second, "Initial wait time")
	flag.DurationVar(&maxWaitTime, "max", time.Second*300, "Max wait time")
	flag.Parse()
	
	u, err := url.ParseRequestURI(targetURL)
	if err != nil || u.Scheme == "" || u.Host == "" {
		log.Fatal("URL Parse Error")
	}

	log.Printf("Start probe HTTP Persistent Connection Time, between %v and %v", initialWaitTime, maxWaitTime)
	probe := httppconntime.NewProbe(initialWaitTime, maxWaitTime, logger)
	pConnTime, err := probe.Do(targetURL)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("HTTP Persistent Connection Time(KeepAlive Time) = %v", pConnTime)
}

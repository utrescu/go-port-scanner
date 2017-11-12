package main

import (
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/utrescu/listIP"
)

var (
	portNumber int
	timeout    string
	debug      *bool
)

func init() {
	flag.IntVar(&portNumber, "port", 22, "Port to scan")
	flag.IntVar(&portNumber, "p", 22, "Port to scan")
	flag.StringVar(&timeout, "timeout", "1000ms", "timeout")
	debug = flag.Bool("v", false, "Show failed connections")
}

func outputFormat(title string, resultats []string) {
	fmt.Println(title)
	fmt.Println("---------------")
	for i := range resultats {
		fmt.Println(resultats[i])
	}
}

func main() {

	var rangs = []string{"192.168.88.0/24"}

	flag.Parse()

	argsWithoutProg := flag.Args()

	if len(argsWithoutProg) > 0 {
		rangs = argsWithoutProg
	}

	_, err := time.ParseDuration(timeout)

	if err != nil {
		log.Fatal("Durada desconeguda\n", err)
	}

	startTime := time.Now()

	resultats, errors := listIP.Check(rangs, portNumber, timeout)

	if *debug {
		outputFormat("errors", errors)
	}
	outputFormat("resultat", resultats)
	scanDuration := time.Since(startTime)
	fmt.Printf("\ndurada: %v\n\n", scanDuration)
}

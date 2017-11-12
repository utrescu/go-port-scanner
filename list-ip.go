package main

import (
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/utrescu/listIP"
)

const maquines = "Màquines"

var (
	portNumber int
	timeout    string
	debug      bool
	header     string
)

func init() {
	flag.IntVar(&portNumber, "port", 22, "Port to scan")
	flag.IntVar(&portNumber, "p", 22, "Port to scan")
	flag.StringVar(&timeout, "timeout", "1000ms", "timeout")
	flag.BoolVar(&debug, "v", false, "Show failed connections")
	flag.StringVar(&header, "h", maquines, "Header Text")
	flag.StringVar(&header, "header", maquines, "Header Text")
}

func outputFormat(title string, resultats []string) {
	fmt.Println(title)
	if header == maquines {
		fmt.Println("--------------------")
	}
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

	if debug {
		outputFormat("errors", errors)
	}
	outputFormat(header, resultats)
	scanDuration := time.Since(startTime)

	if header == maquines {
		fmt.Printf("\ndurada: %v\n\n", scanDuration)
	}
}

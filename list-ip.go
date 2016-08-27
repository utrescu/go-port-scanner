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
)

func init() {
	flag.IntVar(&portNumber, "port", 22, "Port to scan")
	flag.IntVar(&portNumber, "p", 22, "Port to scan")
	flag.StringVar(&timeout, "timeout", "1000ms", "Temps d'espera")
}

func outputFormat(resultats []string, scanDuration time.Duration) {
	fmt.Println("Màquines")
	fmt.Println("---------------")
	for i := range resultats {
		fmt.Println(resultats[i])
	}

	fmt.Printf("\ndurada: %v\n\n", scanDuration)
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

	resultats := listIP.Check(rangs, portNumber, timeout)

	scanDuration := time.Since(startTime)
	outputFormat(resultats, scanDuration)
}

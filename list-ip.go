package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"strconv"
	"sync"
	"time"
)

var wg sync.WaitGroup
var portNumber int
var timeout string

/*
IPList: Obtenir la llista d'IPs
*/
type IPList struct {
	ip    []net.IP
	alive []string
}

/*
  fill: Calcula totes les adreces de la xarxa
*/
func (n *IPList) fill(ip net.IP, ipnet *net.IPNet) {
	notfirst := false
	for ip := ip.Mask(ipnet.Mask); ipnet.Contains(ip); inc(ip) {

		if !notfirst {
			// remove network address
			notfirst = true
			continue
		}
		novaIP := make(net.IP, len(ip))
		copy(novaIP, ip)

		n.ip = append(n.ip, novaIP)

	}
	// Remove broadcast
	n.ip = n.ip[0 : len(n.ip)-1]

}

func inc(ip net.IP) {
	for j := len(ip) - 1; j >= 0; j-- {
		ip[j]++
		if ip[j] > 0 {
			break
		}
	}
}

func rebreResultats(outChan <-chan string, resultChan chan<- []string) {
	var alives []string

	for s := range outChan {
		alives = append(alives, s)
	}

	resultChan <- alives
}

func (n *IPList) comprovaVius(port int) {

	outChan := make(chan string, len(n.ip))
	resultChan := make(chan []string)

	numHosts := len(n.ip)

	fmt.Println("Hosts ", numHosts)

	wg.Add(numHosts)

	for i := 0; i < numHosts; i++ {
		go estaViu(n.ip[i], port, outChan)
	}

	wg.Wait()

	// Per poder fer servir el rang
	close(outChan)

	go rebreResultats(outChan, resultChan)
	n.alive = <-resultChan
}

func estaViu(ip net.IP, port int, outChan chan<- string) {

	defer wg.Done()
	timeoutDuration, err := time.ParseDuration(timeout)

	connexio := ip.String() + ":" + strconv.Itoa(port)

	conn, err := net.DialTimeout("tcp", connexio, timeoutDuration)
	if err == nil {
		fmt.Println("Obert " + ip.String())
		outChan <- ip.String()
		conn.Close()
	}
	return
}

func init() {
	flag.IntVar(&portNumber, "port", 22, "Port to scan")
	flag.IntVar(&portNumber, "p", 22, "Port to scan")
	flag.StringVar(&timeout, "timeout", "1000ms", "Temps d'espera")
}

func main() {
	var listIPs IPList
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

	for rang := range rangs {
		ip, ipnet, err := net.ParseCIDR(rangs[rang])
		if err != nil {
			log.Fatal(err)
		}
		listIPs.fill(ip, ipnet)
	}

	startTime := time.Now()

	listIPs.comprovaVius(portNumber)
	fmt.Println(listIPs.alive)

	scanDuration := time.Since(startTime)
	fmt.Printf("\nFet en %v\n", scanDuration)

}

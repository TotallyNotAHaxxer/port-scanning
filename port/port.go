package port

import (
	"fmt"
	"net"
	"strconv"
	"sync"
	"time"
	
	"github.com/ArkAngeL43/port-scanning/port"

)

func ScanPort(protocol, hostname, service string, port int, resultChannel chan PortResult, wg *sync.WaitGroup) {
	defer wg.Done()
	result := PortResult{Port: port, Service: service}
	address := hostname + ":" + strconv.Itoa(port)
	conn, err := net.DialTimeout(protocol, address, 1*time.Second)
	if err != nil {
		result.State = false
		resultChannel <- result
		return
	}
	defer conn.Close()
	result.State = true
	resultChannel <- result
	return
}

func ScanPorts(hostname string, ports PortRange) (ScanResult, error) {
	var results []PortResult
	var scanned ScanResult
	var wg sync.WaitGroup

	resultChannel := make(chan PortResult, ports.End-ports.Start)

	addr, err := net.LookupIP(hostname)
	if err != nil {
		fmt.Println("\033[31m[-] Error occured during scanning.....")
		return scanned, err
	}
	for i := ports.Start; i <= ports.End; i++ {
		if service, ok := common[i]; ok {
			wg.Add(1)
			go ScanPort("tcp", hostname, service, i, resultChannel, &wg)

		}
	}
	wg.Wait()
	close(resultChannel)
	for result := range resultChannel {
		results = append(results, result)
	}

	scanned = ScanResult{
		hostname: hostname,
		ip:       addr,
		results:  results,
	}
	return scanned, nil
}

func DisplayScanResult(result ScanResult) {
	ip := result.ip[len(result.ip)-1]
	fmt.Printf("\n\t\033[34m[\033[35m*\033[34m] Scan Results for   ├ %s (%s)\n", result.hostname, ip.String())
	for _, v := range result.results {
		if v.State {
			fmt.Printf("\t\033[34m[\033[35m+\033[34m]\t\t\t┡ %d	%s\n", v.Port, v.Service)
		}
	}
}

func GetOpenPorts(hostname string, ports PortRange) {
	scanned, err := ScanPorts(hostname, ports)
	checkErr(err)
	DisplayScanResult(scanned)
}

/*
Copyright (c) 2022 Andrew Repp auguste@andrewrepp.com

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/
package ports

import (
	"fmt"
	"net"
	"strconv"
	"time"
)

type ScanJob struct {
	Port int
	Protocol string
}

type ScanResult struct {
	Port int
	State string
	Protocol string
}

func ToSlice(c chan ScanResult) []ScanResult {
    s := make([]ScanResult, 0)
    for i := range c {
        s = append(s, i)
    }
    return s
}

func ScanPort(protocol, hostname string, port int, results chan ScanResult)  {
	result := ScanResult{Port: port, Protocol: protocol}
	address := hostname + ":" + strconv.Itoa(port)
	conn, err := net.DialTimeout(protocol, address, 60*time.Second)

	if err != nil {
		result.State = "Closed"
		results <- result
	}
	defer conn.Close()
	result.State = "Open"
	results <- result
}

func ScanEngine(scanjobs []ScanJob, target string) []ScanResult{
	results_channel :=make(chan ScanResult)
	defer close(results_channel)

	for x := 0; x < len(scanjobs); x++ {
		go ScanPort(scanjobs[x].Protocol, target, scanjobs[x].Port, results_channel)
	}
	fmt.Println("AJR got past calling loop")
	showme := <- results_channel
	fmt.Println(showme)
	scan_results := []ScanResult{}
	return scan_results
}
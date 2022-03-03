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
package main

import (
	"flag"
	"fmt"

	"github.com/ReppCodes/auguste/ports"
)

func main() {
	// CLI handling
	var scan_type string
	var in_ports string
	var target string = "nmap.scanme.org" // TODO add cli flag to indicate target
	flag.StringVar(&scan_type, "t", "none", "Specify type of scan. Default is tcp active scan")
	// TODO -- scan types to support.  figure out how to list in help function
	// Ping Scan - This sends a ping and listens for a response
	// TCP Half-Open - Also known as SYN scan, these scans attempt to start a TCP connection, listen for the SYN-ACK response and then never send the final ACK.
	// TCP Open - This is just attempting to open a TCP connection on a host:port like we have done above
	// UDP - Very similar to TCP scanning except using the UDP protocol.
	// Stealth Scanning - A far more sophisticated type of scan which has been designed so that these scans don’t show up in connection logs.
	// Service Probes -- for each of the above?
	flag.StringVar(&in_ports, "p", "none", "Specify ports to scan, format is either 0 or 0-65535. Default behavior is most commonly used ports")
	flag.Parse()

	// parse scan type
	if scan_type == "none" {
		fmt.Println("no scan_type provided yo")
	}

	// parse ports if provided
	if in_ports == "none" {
		fmt.Println("No ports indicated, scanning range of most commonly used ports")
		var common_ports = ports.Get_common_ports()
		var scan_jobs = []ports.ScanJob{}
		for x := 0; x < len(common_ports); x++ {
			var new_job = ports.ScanJob{Port: common_ports[x]}
			scan_jobs = append(scan_jobs, new_job)
		}
		results := ports.ScanEngine(scan_jobs, target)
		for _, value := range results {
			fmt.Printf("%+v", value)
		}
	} else {
		fmt.Printf("Port argument provided: %s\n", in_ports)
	}
}

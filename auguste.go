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
)

func main() {
	// CLI handling
	var scan_type string
	var ports string
	flag.StringVar(&scan_type, "t", "none", "Specify type of scan. Default is tcp active scan")
	flag.StringVar(&ports, "p", "none", "Specify ports to scan. Default behavior is most commonly used ports")
	flag.Parse()

	// parse ports if provided
	if ports == "none" {
		fmt.Println("no ports provided yo")
		var common_ports = get_common_ports()
		for x := 0; x < len(get_common_ports()); x++ {
			fmt.Printf("%d\n", common_ports[x])
		}
	}

	// parse scan type
	if scan_type == "none" {
		fmt.Println("no scan_type provided yo")
	}

}

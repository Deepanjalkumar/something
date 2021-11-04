package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

func Extract_urls(domain string, output string) {

	url_byte, _ := ioutil.ReadAll(r.Body)

	data := string(url_byte)

	filename, _ := os.Create(output)
	filename.WriteString(fmt.Sprintf("%s\n", data))
}

func main() {

	domain := flag.String("domain", "", "Enter the domain name")

	output := flag.String("output", "", "Enter the output filename")

	flag.Parse()

	Extract_urls(*domain, *output)
}

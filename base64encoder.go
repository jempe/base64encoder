package main

import (
	"flag"
	b64 "encoding/base64"
	"fmt"
	"io/ioutil"
)

var file  = flag.String("file", "", "Define what file you want to encode")

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main() {
	flag.Parse()
	if *file != "" {
		dat, err := ioutil.ReadFile(*file)
		check(err)
		
		sEnc := b64.StdEncoding.EncodeToString(dat)
		fmt.Println(sEnc)
	} else {
		fmt.Print("Please select a file")
	}
}
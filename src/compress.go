package main

import (
	"flag"
	"log"
	"io/ioutil"
	"gopkg.in/kothar/brotli-go.v0/enc"
	"os"
)

var file = flag.String("file", "", "file to compress")

func main() {
	flag.Parse()	
	if *file == "" {
		log.Fatal("You didn't pass a file")
	}
	content, err := ioutil.ReadFile(*file)
	if err != nil {
		log.Fatal("Cannot read file")
	}
	compressedData, err := enc.CompressBuffer(nil, content, make([]byte, 0))
	if err != nil {
		log.Fatal("Cannot compress file")
	}
	newFile, err := os.Create(*file+".br")
	if err != nil {
		log.Fatal("Cannot create new file")
	}
	defer newFile.Close()
	_, err := newFile.Write(compressedData)
	if err != nil {
		log.Fatal("Cannot write to a file")
	}
}


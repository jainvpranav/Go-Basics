package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func readFileName() string {
	return os.Args[1]
}

func readFile() []byte {
	file, err := os.ReadFile(readFileName())
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	return file
}

func openFile() {
	f, err := os.Open(readFileName())
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(0)
	}
	io.Copy(os.Stdout, f)
}

func readFileFromFileName() {
	fmt.Println(string(readFile()))
	openFile()
}

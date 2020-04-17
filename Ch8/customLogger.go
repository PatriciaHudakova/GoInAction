package main

import (
	"io"
	"io/ioutil"
	"log"
	"os"
)

var (
	Trace   *log.Logger //miscellaneous
	Info    *log.Logger //important
	Warning *log.Logger //be concerned
	Error   *log.Logger //critical problem
)

func init() {
	file, err := os.OpenFile("errors.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("failed to open", err) //not explained in the book?
	}

	//destination.prefix/flags
	Trace = log.New(ioutil.Discard, "Trace: ", log.Ldate|log.Ltime|log.Lshortfile)

	Info = log.New(os.Stdout, "Info: ", log.Ldate|log.Ltime|log.Lshortfile)

	Warning = log.New(os.Stdout, "Warning: ", log.Ldate|log.Ltime|log.Lshortfile)

	Error = log.New(io.MultiWriter(file, os.Stderr), "Error: ", log.Ldate|log.Ltime|log.Lshortfile)
}

func main() {
	Trace.Println("I have something standard to say")
	Info.Println("Special Information")
	Warning.Println("There is something you need to know about")
	Error.Println("Something has failed")
}

package main

import "log"

func init() {
	log.SetPrefix("TRACE: ")
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)
}

func main() {
	log.Println("message")
	log.Fatalln("fatal message") // call to os.Exit(1)
	log.Panicln("panic message") //call to panic()
}

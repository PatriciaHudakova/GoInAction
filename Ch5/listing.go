package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

func main() {
	var b bytes.Buffer
	b.Write([]byte("Hello")) //write a string to the buffer
	fmt.Fprint(&b, "World")  //concatenate string to the buffer
	io.Copy(os.Stdout, &b)   //buffer ==> stdout
}

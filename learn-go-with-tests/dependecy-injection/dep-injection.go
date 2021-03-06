package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
)

func Greet1(name string) {
	fmt.Printf("Hello, %s", name)
}

func Greet2(writer *bytes.Buffer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}

func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}

func MyGreeterHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "world")
}

func main() {
	log.Fatal(http.ListenAndServe(":5000", http.HandlerFunc(MyGreeterHandler)))
}

// PRINTF CODE
// Printf formats according to a format specifier and writes to standard output.
// It returns the number of bytes written and any write error encountered.
// func Printf(format string, a ...interface{}) (n int, err error) {
// 	return Fprintf(os.Stdout, format, a...)
// }

// CODE Fprintf
//func Fprintf(w io.Writer, format string, a ...interface{}) (n int, err error) {
//	p := newPrinter()
//	p.doPrintf(format, a)
//	n, err = w.Write(p.buf)
//  p.free()
//  return
//}

// TYPE WRITER
// type Writer interface {
//    Write(p []byte) (n int, err error)
// }

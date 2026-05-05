package main

import (
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	// first exercise - Creates a strings.NewReader("hello, go reader"), then copies its contents to os.Stdout using io.Copy
	r := strings.NewReader("hello, go reader\n\n")
	io.Copy(os.Stdout, r)

	// second exercise - Reads a compressed file
	f, err := os.Open("exercises/week-09/day49/myfile.gz")
	if err != nil {
		fmt.Printf("Unable to read file %v\n\n", err)
		return
	}
	gzipReader, err := gzip.NewReader(f)
	if err != nil {
		fmt.Printf("Unable to create gzip reader %v\n\n", err)
		return
	}
	io.Copy(os.Stdout, gzipReader)
	gzipReader.Close()
	f.Close()

	// third exercise - creates a compressed file
	f, err = os.Create("exercises/week-09/day49/test.gz")
	if err != nil {
		fmt.Printf("Unable to create gzip writer %v\n\n", err)
		return
	}
	gzipWriter := gzip.NewWriter(f)

	// there are different ways to write to a writer
	gzipWriter.Write([]byte("from-gzip-writer.write\n\n"))
	fmt.Fprintf(gzipWriter, "from-fmt-fprintf\n\n")
	io.WriteString(gzipWriter, "from-io-writestring\n\n")

	gzipWriter.Close()
	f.Close()

	// fourth exercise - reads from compressed file that was created in third exercise
	f, err = os.Open("exercises/week-09/day49/test.gz")
	if err != nil {
		fmt.Printf("Unable to read file %v\n\n", err)
		return
	}
	gzipReader, err = gzip.NewReader(f)
	if err != nil {
		fmt.Printf("Unable to create gzip reader %v\n\n", err)
		return
	}
	io.Copy(os.Stdout, gzipReader)
	gzipReader.Close()
	f.Close()
}

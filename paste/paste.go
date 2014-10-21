// Paste from standard in or file to the server
// TODO: help messages; fix buffer size; read port and host from flags;
//		multiple file read
package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

const port = ":8080"
const host = "localhost"

func main() {
	var buf bytes.Buffer
	if len(os.Args) > 1 {
		file, err := os.Open(os.Args[1])
		if err != nil {
			log.Fatal(err)
		}
		b := make([]byte, 100) //FIXME
		file.Read(b)
		buf.Write([]byte(b))
	} else {
		b, err := ioutil.ReadAll(os.Stdin)
		if err != nil {
			log.Fatal("Can't read from STDIN")
		}
		buf.Write(b)
	}
	resp, err := http.Post("http://"+host+port, "text", &buf)
	if err != nil {
		log.Fatal(err)
	}
	content, _ := ioutil.ReadAll(resp.Body)
	fmt.Printf("%s", content)
}
